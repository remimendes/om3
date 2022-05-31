package daemonps

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"opensvc.com/opensvc/core/event"
	ps "opensvc.com/opensvc/util/pubsub"
)

func TestDaemonPubSub(t *testing.T) {
	bus := &ps.T{}
	cmdC, err := bus.Start(context.Background(), t.Name())
	require.Nil(t, err)
	defer bus.Stop()
	var (
		eventKinds    = []string{"hb_stale", "hb_beating"}
		expectedKinds = []string{"event-subscribe", "hb_stale", "hb_beating"}
		detectedKinds []string
	)
	defer UnSubEvent(
		cmdC,
		SubEvent(cmdC,
			"description 1",
			func(e event.Event) {
				t.Logf("detected event %s", e.Kind)
				detectedKinds = append(detectedKinds, e.Kind)
			}))
	time.Sleep(1 * time.Millisecond)
	for _, kind := range eventKinds {
		PubEvent(cmdC, event.Event{Kind: kind})
	}
	time.Sleep(1 * time.Millisecond)
	require.ElementsMatch(t, expectedKinds, detectedKinds)
}

func TestNamespacesAreDeclared(t *testing.T) {
	_ = NsAll
	_ = NsCfg
	_ = NsCfgFile
	_ = NsStatus
	_ = NsSmon
	_ = NsSetSmon
	_ = NsAgg
}
