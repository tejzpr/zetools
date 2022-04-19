package ping

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"zetools/commands"

	"github.com/go-ping/ping"
	"github.com/urfave/cli/v2"
)

func init() {
	commands.RegisterCommand(PingCommandName, &pingCommand{}, nil)
}

// PingCommandName returns the name of the command
const PingCommandName commands.CommandName = "ping"

// pingCommand is the ping command
type pingCommand struct {
}

// Name returns the name of the command
func (b *pingCommand) Name() commands.CommandName {
	return PingCommandName
}

// Aliases returns the aliases of the command
func (b *pingCommand) Aliases() []string {
	return []string{"pn"}
}

// Usage returns the usage of the command
func (b *pingCommand) Usage() string {
	return "Ping"
}

// Subcommands returns the subcommands of the command
func (b *pingCommand) Subcommands() []*cli.Command {
	return []*cli.Command{}
}

// Flags returns the flags of the command
func (b *pingCommand) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "tries",
			Aliases:  []string{"t"},
			Value:    "3",
			Usage:    "Number of tries",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "host",
			Value:    "",
			Usage:    "Host to ping",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "timeout",
			Aliases:  []string{"to"},
			Value:    "5s",
			Usage:    "Connection timeout",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "privileged",
			Aliases:  []string{"p"},
			Value:    "true",
			Usage:    "Run as privileged (root), required to ping some hosts",
			Required: false,
		},
	}
}

// Action returns the action of the command
func (b *pingCommand) Action(c *cli.Context) error {
	tries := c.Int("tries")
	host := c.String("host")
	timeout := c.Duration("timeout")
	privileged := c.Bool("privileged")
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return err
	}
	pinger.Interval = time.Second
	pinger.SetPrivileged(privileged)
	pinger.Timeout = timeout
	pinger.Count = tries
	// Listen for Ctrl-C.
	cos := make(chan os.Signal, 1)
	signal.Notify(cos, os.Interrupt)
	go func() {
		for _ = range cos {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		return err
	}
	return nil
}
