package domain

import "testing"

func TestPermission_Has(t *testing.T) {
	type args struct {
		permission Permission
	}
	tests := []struct {
		name string
		p    Permission
		args args
		want bool
	}{
		{
			name: "Verified has NeedVerified",
			p:    IsVerified,
			args: args{
				permission: NeedVerified,
			},
			want: true,
		},
		{
			name: "Moderator has NeedVerified",
			p:    IsModerator,
			args: args{permission: NeedVerified},
			want: true,
		},
		{
			name: "Admin has NeedVerified",
			p:    IsAdmin,
			args: args{permission: NeedVerified},
			want: true,
		},
		{
			name: "Moderator has NeedModerator",
			p:    IsModerator,
			args: args{permission: NeedModerator},
			want: true,
		},
		{
			name: "Admin has NeedModerator",
			p:    IsAdmin,
			args: args{permission: NeedModerator},
			want: true,
		},
		{
			name: "Admin has NeedAdmin",
			p:    IsAdmin,
			args: args{permission: NeedAdmin},
			want: true,
		},
		{
			name: "Unknown doesn't have NeedVerified",
			p:    IsUnknown,
			args: args{permission: NeedVerified},
			want: false,
		},
		{
			name: "Unknown doesn't have NeedModerator",
			p:    IsUnknown,
			args: args{permission: NeedModerator},
			want: false,
		},
		{
			name: "Unknown doesn't have NeedAdmin",
			p:    IsUnknown,
			args: args{permission: NeedAdmin},
			want: false,
		},
		{
			name: "Verified doesn't have NeedModerator",
			p:    IsVerified,
			args: args{permission: NeedModerator},
			want: false,
		},
		{
			name: "Verified doesn't have NeedAdmin",
			p:    IsVerified,
			args: args{permission: NeedAdmin},
			want: false,
		},
		{
			name: "Moderator doesn't have NeedAdmin",
			p:    IsModerator,
			args: args{permission: NeedAdmin},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Has(tt.args.permission); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}
