package tracer

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func InitJaeger() (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := &config.Configuration{
		ServiceName: "assignment",
		// 采样
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	tracer, closer, err = cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	// 设置tracer为全局单例对象
	opentracing.SetGlobalTracer(tracer)
	return
}
