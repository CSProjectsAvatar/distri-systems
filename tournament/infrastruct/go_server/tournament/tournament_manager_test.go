package main

import (
	"strings"
	"testing"
)

func Test_RunCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		result  string
		wantErr bool
	}{
		{
			name: "test_run_command",
			args: args{
				command: "echo hello",
			},
			result:  "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err, out, _ := RunCommand(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("RunCommand() error = %v, wantErr %v", err, tt.wantErr)
			} else if !strings.HasPrefix(out, tt.result) {
				t.Errorf("RunCommand() out = %v, want %v", out, tt.result)
			}
		})
	}
}

func Test_SaveTournament(t *testing.T) {
	type args struct {
		tour_name string
		files     map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_save_tournament",
			args: args{
				tour_name: "tournament_1",
				files: map[string]string{
					"game.py":     "print('1')",
					"player_1.py": "print('1')",
					"player_2.py": "print('2')",
					"player_3.py": "print('3')",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveTournament(tt.args.tour_name, tt.args.files); (err != nil) != tt.wantErr {
				t.Errorf("SaveTournament() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_RunMatch(t *testing.T) {
	type args struct {
		tourN_name string
		player_i   string
		player_j   string
	}
	tests := []struct {
		name     string
		args     args
		expected GameResult
		wantErr  bool
	}{
		{
			name: "test_run_match",
			args: args{
				tourN_name: "tournament_1",
				player_i:   "player_1",
				player_j:   "player_2",
			},
			expected: 1,
			wantErr:  false,
		},
		{
			name: "test_run_match",
			args: args{
				tourN_name: "tournament_1",
				player_i:   "player_2",
				player_j:   "player_3",
			},
			expected: 1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gameR, err := RunMatch(tt.args.tourN_name, tt.args.player_i, tt.args.player_j); (err != nil) != tt.wantErr {
				t.Errorf("RunMatch() error = %v, wantErr %v", err, tt.wantErr)
			} else if gameR != tt.expected {
				t.Errorf("RunMatch() gameR = %v, want %v", gameR, tt.expected)
			}
		})
	}

}
