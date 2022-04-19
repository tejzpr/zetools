package hmac

import "testing"

func TestHMAC(t *testing.T) {
	type args struct {
		text     string
		fileName string
		key      string
		hashType string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "HMAC256-String",
			args: args{text: "test", fileName: "", key: "abc", hashType: "sha256"},
			want: "d64ccf0d4b1449153d78215f9ff9b90ec3730de1fd2b357e153026c9a3fada96",
		},
		{
			name: "HMAC512-String",
			args: args{text: "test", fileName: "", key: "abc", hashType: "sha512"},
			want: "c55194bd4959a626fb2f2f18b6e84a870b3e19f479b57d8837761cc90819592e70d2cb5031f91108f14e8afb01eb42fae785f248c47fab9302c1244fba57971d",
		},
		{
			name: "HMAC256-File",
			args: args{text: "", fileName: "./testdata/hmac/test.txt", key: "abc", hashType: "sha256"},
			want: "99f8aedf2162d644dd683c9b4ecb04d1d4efc2be1fd634eee7856f369b2b90be",
		},
		{
			name: "HMAC512-File",
			args: args{text: "", fileName: "./testdata/hmac/test.txt", key: "abc", hashType: "sha512"},
			want: "dbe37866e0a483bc53f34497056ebbd1ab43e19fa5eb551399ce6c1289c8627842f8c2aed44352c95935533a27a485ff0b3f532c0f6c914d704f087ea1a18b30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HMAC(tt.args.text, tt.args.fileName, tt.args.key, tt.args.hashType)
			if (err != nil) != tt.wantErr {
				t.Errorf("HMAC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HMAC() = %v, want %v", got, tt.want)
			}
		})
	}
}
