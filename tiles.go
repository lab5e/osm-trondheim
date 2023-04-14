package osmtrondheim

import "embed"

//go:embed map/tiles
var tiles embed.FS

// TrondheimOSMTiles returns an embeddded file system with Open Street Map tiles containting data for
// Trondheim and surrounding municipalities
func TrondheimOSMTiles() embed.FS {
	return tiles
}
