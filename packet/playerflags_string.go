// Code generated by "gcraft_stringer -type=PlayerFlags -method=toString"; DO NOT EDIT.

package packet

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PLAYER_FLAGS_NONE-0]
	_ = x[PLAYER_FLAGS_GROUP_LEADER-1]
	_ = x[PLAYER_FLAGS_AFK-2]
	_ = x[PLAYER_FLAGS_DND-4]
	_ = x[PLAYER_FLAGS_GM-8]
	_ = x[PLAYER_FLAGS_GHOST-16]
	_ = x[PLAYER_FLAGS_RESTING-32]
	_ = x[PLAYER_FLAGS_UNK7-64]
	_ = x[PLAYER_FLAGS_FFA_PVP-128]
	_ = x[PLAYER_FLAGS_CONTESTED_PVP-256]
	_ = x[PLAYER_FLAGS_PVP_DESIRED-512]
	_ = x[PLAYER_FLAGS_HIDE_HELM-1024]
	_ = x[PLAYER_FLAGS_HIDE_CLOAK-2048]
	_ = x[PLAYER_FLAGS_PARTIAL_PLAY_TIME-4096]
	_ = x[PLAYER_FLAGS_NO_PLAY_TIME-8192]
	_ = x[PLAYER_FLAGS_UNK15-16384]
	_ = x[PLAYER_FLAGS_UNK16-32768]
	_ = x[PLAYER_FLAGS_SANCTUARY-65536]
	_ = x[PLAYER_FLAGS_TAXI_BENCHMARK-131072]
	_ = x[PLAYER_FLAGS_PVP_TIMER-262144]
}

const _PlayerFlags_name = "PLAYER_FLAGS_NONEPLAYER_FLAGS_GROUP_LEADERPLAYER_FLAGS_AFKPLAYER_FLAGS_DNDPLAYER_FLAGS_GMPLAYER_FLAGS_GHOSTPLAYER_FLAGS_RESTINGPLAYER_FLAGS_UNK7PLAYER_FLAGS_FFA_PVPPLAYER_FLAGS_CONTESTED_PVPPLAYER_FLAGS_PVP_DESIREDPLAYER_FLAGS_HIDE_HELMPLAYER_FLAGS_HIDE_CLOAKPLAYER_FLAGS_PARTIAL_PLAY_TIMEPLAYER_FLAGS_NO_PLAY_TIMEPLAYER_FLAGS_UNK15PLAYER_FLAGS_UNK16PLAYER_FLAGS_SANCTUARYPLAYER_FLAGS_TAXI_BENCHMARKPLAYER_FLAGS_PVP_TIMER"

var _PlayerFlags_map = map[PlayerFlags]string{
	0:      _PlayerFlags_name[0:17],
	1:      _PlayerFlags_name[17:42],
	2:      _PlayerFlags_name[42:58],
	4:      _PlayerFlags_name[58:74],
	8:      _PlayerFlags_name[74:89],
	16:     _PlayerFlags_name[89:107],
	32:     _PlayerFlags_name[107:127],
	64:     _PlayerFlags_name[127:144],
	128:    _PlayerFlags_name[144:164],
	256:    _PlayerFlags_name[164:190],
	512:    _PlayerFlags_name[190:214],
	1024:   _PlayerFlags_name[214:236],
	2048:   _PlayerFlags_name[236:259],
	4096:   _PlayerFlags_name[259:289],
	8192:   _PlayerFlags_name[289:314],
	16384:  _PlayerFlags_name[314:332],
	32768:  _PlayerFlags_name[332:350],
	65536:  _PlayerFlags_name[350:372],
	131072: _PlayerFlags_name[372:399],
	262144: _PlayerFlags_name[399:421],
}

func (i PlayerFlags) toString() string {
	if str, ok := _PlayerFlags_map[i]; ok {
		return str
	}
	return "PlayerFlags(" + strconv.FormatInt(int64(i), 10) + ")"
}
