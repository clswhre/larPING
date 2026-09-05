package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "embed"

	"github.com/likexian/whois"
	probing "github.com/prometheus-community/pro-bing"
	flag "github.com/spf13/pflag"
)

type Config struct {
	Target string
	IP     net.IP
	Count  int
	Whois  bool
	Size   int
}

//go:embed logo.txt
var asciiLogo string

func main() {

	cfg := parseArgs()

	printTypewriter(1*time.Millisecond, colorBlue, "%s\n", asciiLogo)
	printTypewriter(50*time.Millisecond, colorGreen, "%s", startingText)
	printTypewriter(50*time.Millisecond, colorRed, "%s\n", cfg.Target)

	time.Sleep(300 * time.Millisecond)

	checkIP(cfg)
	if cfg.Whois {
		whoisQuery(cfg.Target)
	}

	printTypewriter(600*time.Millisecond, colorBlue, "...\n")
	pingTarget(cfg)
}

func parseArgs() Config {
	var cfg Config

	flag.IntVarP(&cfg.Count, "count", "c", 0, "Numb3r 0f p4ck37s 70 s3nd (0 = infini73)")
	flag.BoolVarP(&cfg.Whois, "whois", "w", false, "P3rf0rm WH0IS l00kup 0n 74rg37")
	flag.IntVarP(&cfg.Size, "bytes", "b", 56, "ICMP p4ck37 p4yl04d siz3 in by73s")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "US4G3: larPING [0pt10n5] <ip_or_host>\n\n0pt10ns:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("[!] 3RR0R! M41NFR4M3 0V3RL04D!")
		flag.Usage()
		os.Exit(1)
	}

	cfg.Target = args[0]
	parsedIP := net.ParseIP(cfg.Target)
	if parsedIP == nil {

		ips, err := net.LookupIP(cfg.Target)
		if err != nil || len(ips) == 0 {
			fmt.Println("[!] 3RR0R! UN48L3 T0 R350LV3 T4RG3T :c")
			flag.Usage()
			os.Exit(1)
		}
		parsedIP = ips[0]
	}
	cfg.IP = parsedIP

	return cfg
}

func pingTarget(cfg Config) {
	pinger, err := probing.NewPinger(cfg.IP.String())
	if err != nil {
		printColor(colorRed, errorText, err)
		return
	}

	pinger.SetPrivileged(false)
	pinger.Size = cfg.Size
	if cfg.Count > 0 {
		pinger.Count = cfg.Count
	}

	pinger.OnRecv = func(pkt *probing.Packet) {
		printColor(colorBlue, packetText, pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt.Round(time.Millisecond))
		printColor(colorYellow, "%s", phrases[rand.IntN(len(phrases))])

		if rand.IntN(4) == 0 {
			printLarpLogos()
		}
		printColor(colorCyan, "----------------------------------------")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		pinger.Stop()
	}()

	printColor(colorYellow, initConnection, cfg.Target)

	if err := pinger.Run(); err != nil {
		printColor(colorRed, failText, err)
		return
	}

	stats := pinger.Statistics()
	printColor(colorCyan, headerText, stats.Addr)
	printColor(colorWhite, statText, stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	fmt.Println("[+] M1SS10N C0MPL3T3. 3X1T...")
}

func whoisQuery(target string) {
	printTypewriter(50*time.Millisecond, colorYellow, "%s\n", whoisText)
	time.Sleep(300 * time.Millisecond)

	result, err := whois.Whois(target)
	printColor(colorWhite, "\n----------")
	if err != nil {
		printColor(colorRed, whoisErrText, err)
		return
	}

	printColor(colorBlue, "%s", result)
}

func checkIP(cfg Config) {
	if cfg.IP.To16().String() == "127.0.0.1" {
		printTypewriter(50*time.Millisecond, colorRed, rmRfWarningText)
		printColor(colorYellow, "\n[?] JUS7 K1DD1NG... 0R 4M I? ;)\n")
		os.Exit(2)
	}
	if cfg.IP.To16().String() == "1.1.1.1" {
		printTypewriter(100*time.Millisecond, colorRed, "Okay... Whatever\n")
	}
}
