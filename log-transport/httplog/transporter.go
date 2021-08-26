package httplog

import (
	"fmt"

	log_transport "github.com/eolinker/goku/log-transport"

	"github.com/eolinker/eosc"
	eosc_log "github.com/eolinker/eosc/log"
)

type Transporter struct {
	*eosc_log.Transporter
	writer *_HttpWriter
}

func (t *Transporter) Reset(c interface{}, formatter eosc_log.Formatter) error {
	conf, ok := c.(*Config)
	if !ok {
		return fmt.Errorf("need %s,now %s", eosc.TypeNameOf((*Config)(nil)), eosc.TypeNameOf(c))
	}

	t.Transporter.SetFormatter(formatter)
	return t.reset(conf)
}

func (t *Transporter) Close() error {
	t.Transporter.Close()
	return t.writer.Close()
}

func (t *Transporter) reset(c *Config) error {
	t.SetLevel(c.Level)

	t.writer.reset(c)
	t.Transporter.SetOutput(t.writer)
	return nil
}

func CreateTransporter(conf *Config, formatter eosc_log.Formatter) (log_transport.TransporterReset, error) {

	httpWriter := newHttpWriter()

	transport := &Transporter{
		Transporter: eosc_log.NewTransport(httpWriter, conf.Level, formatter),
		writer:      httpWriter,
	}

	e := transport.Reset(conf, formatter)
	if e != nil {
		return nil, e
	}
	return transport, nil
}
