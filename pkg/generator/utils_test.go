package generator

import (
	"testing"

	"github.com/ElfLabs/id-generator/pkg/format"
	"github.com/ElfLabs/id-generator/pkg/planner/mid"
)

func TestParseID(t *testing.T) {
	type args struct {
		id        ID
		formatter Formatter
	}
	tests := []struct {
		name       string
		args       args
		wantRegion int64
		wantNode   int64
		wantCount  int64
		wantStep   int64
	}{
		{
			name: "snowflake",
			args: args{
				id:        1698273118818144257,
				formatter: format.NewSnowflakeFormat(),
			},
			wantRegion: 0,
			wantNode:   1,
			wantCount:  404899863915,
			wantStep:   1,
		},
		{
			name: "mid",
			args: args{
				id:        1000001686,
				formatter: mid.NewMidPlanner(mid.DefaultStartMid),
			},
			wantRegion: 2,
			wantNode:   2,
			wantCount:  3906256,
			wantStep:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRegion, gotNode, gotCount, gotStep := ParseID(tt.args.id, tt.args.formatter)
			if gotRegion != tt.wantRegion {
				t.Errorf("ParseID() gotRegion = %v, want %v", gotRegion, tt.wantRegion)
			}
			if gotNode != tt.wantNode {
				t.Errorf("ParseID() gotNode = %v, want %v", gotNode, tt.wantNode)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ParseID() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
			if gotStep != tt.wantStep {
				t.Errorf("ParseID() gotStep = %v, want %v", gotStep, tt.wantStep)
			}
		})
	}
}
