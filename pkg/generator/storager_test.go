package generator

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ElfLabs/id-generator/pkg/planner/mid"
	"github.com/ElfLabs/id-generator/pkg/planner/snowflake"
	"github.com/ElfLabs/id-generator/pkg/storager/file"
)

func TestStoragerGenerator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type args struct {
		region  int64
		node    int64
		options []Option
		times   int
	}
	tests := []struct {
		name  string
		args  args
		check func(args args, ids []ID) error
	}{
		{
			name: "snowflake",
			args: args{
				node: 1,
				options: []Option{
					WithEpoch(snowflake.DefaultEpoch),
					WithStopChan(ctx.Done()),
					WithBuffer(1024),
					WithStorager(file.NewFileStorage(filepath.Join(os.TempDir(), "snowflake.json"))),
				},
				times: 4096 * 10,
			},
			check: func(args args, ids []ID) error {
				for _, id := range ids {
					if count(ids, id) > 1 {
						return fmt.Errorf("%d is repetition", id)
					}
				}
				t.Logf(" start: %d end: %d", ids[0], ids[len(ids)-1])
				return nil
			},
		},
		{
			name: "mid",
			args: args{
				region: 2,
				node:   1,
				options: []Option{
					WithStopChan(ctx.Done()),
					WithBuffer(1024),
					WithPlanner(mid.NewMidPlanner(mid.DefaultStartMid)),
					WithStorager(file.NewFileStorage(filepath.Join(os.TempDir(), "mid.json"))),
				},
				times: 4096 * 10,
			},
			check: func(args args, ids []ID) error {
				for _, id := range ids {
					if count(ids, id) > 1 {
						return fmt.Errorf("%d is repetition", id)
					}
				}
				t.Logf(" start: %d end: %d", ids[0], ids[len(ids)-1])
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewGenerator(tt.args.region, tt.args.node, tt.args.options...)
			if err != nil {
				t.Errorf("NewGenerator failed: %v", err)
				return
			}
			var ids = make([]ID, 0, tt.args.times)
			for i := 0; i < tt.args.times; i++ {
				id := g.Generate()
				ids = append(ids, id)
			}
			if tt.check != nil {
				err = tt.check(tt.args, ids)
				if err != nil {
					t.Errorf("Check Generate failed: %v", err)
				}
			}
		})
	}
}

func TestStoragerGeneratorSequencer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	os.RemoveAll(filepath.Join(os.TempDir(), "mid-0-100.json"))
	os.RemoveAll(filepath.Join(os.TempDir(), "mid-20.json"))

	region := int64(2)
	node := int64(1)
	times := 4096 * 10
	seq_0_100 := []Option{
		WithStopChan(ctx.Done()),
		WithBuffer(1024),
		WithPlanner(mid.NewMidPlanner(mid.DefaultStartMid)),
		WithStorager(file.NewFileStorage(filepath.Join(os.TempDir(), "mid-0-100.json"))),
	}

	g_0_100, err := NewGenerator(region, node, seq_0_100...)
	if err != nil {
		t.Errorf("NewGenerator failed: %v", err)
		return
	}

	var ids_0_100 = make([]ID, 0, times)
	for i := 0; i < times; i++ {
		id := g_0_100.Generate()
		ids_0_100 = append(ids_0_100, id)
	}

	var ids_20 []ID
	for {
		seq_20 := []Option{
			WithStopChan(ctx.Done()),
			WithBuffer(1024),
			WithPlanner(mid.NewMidPlanner(mid.DefaultStartMid)),
			WithStorager(file.NewFileStorage(filepath.Join(os.TempDir(), "mid-20.json"))),
		}

		g_20, err := NewGenerator(region, node, seq_20...)
		if err != nil {
			t.Errorf("NewGenerator failed: %v", err)
			return
		}
		id := g_20.Generate()
		ids_20 = append(ids_20, id)
		if len(ids_20) == len(ids_0_100) {
			break
		}
	}

	for idx, id := range ids_0_100 {
		if ids_20[idx] != id {
			t.Errorf("got %d want %d", ids_20[idx], id)
			t.Failed()
		}
	}

}
