package models

import (
	"encoding/json"
	"time"
)

type ResponseRoot struct {
	Data    []VariableResource
	Errors  []Error
	Href    string
	Meta    ResponseRootMeta
	Next    string
	Results ResponseRootResults
}

type ResponseRootMeta struct {
}

type Error struct {
	Code   string
	Detail string
	Id     string
	Error  ErrorSource
	Status string
	Title  string
}

type ErrorSource struct {
	Parameter string
	Pointer   string //JSONPointer
}

type ResponseRootResults struct {
}

type ResourceMeta struct {
}

type ResourceAttributes struct {
}

type ResourceRelationships struct {
	Data []VariableResource
	Href string
	Meta RelationshipMeta
	Next string
}

type RelationshipMeta struct {
}

type Artwork struct {
	BgColor    string
	Height     int
	Width      int
	TextColor1 string
	TextColor2 string
	TextColor3 string
	TextColor4 string
	Url        string
}

type EditorialNotes struct {
	Short    string
	Standard string
}

type PlayParameters struct {
	Id        string
	Kind      string
	IsLibrary bool //UnDocumented
}

type Preview struct {
	Artwork Artwork
	Url     string
}

type Relationship struct {
}

type TrackRelationship struct {
	Relationship
	Data []SongOrMusicVideoResource //(Song | MusicVideo)] (Required) The data for the track included in the relationship.
}

type VariableResponse struct {
	Data []VariableResource
}
type SongOrMusicVideoResource struct {
	Type       string
	Song       *Song
	MusicVideo *MusicVideo
}
type VariableResource struct {
	Type          string
	Curator       *Curator
	AppleCurator  *AppleCurator
	Artist        *Artist
	LibraryArtist *LibraryArtist
	Rating        *Rating
	Station       *Station
	Storefront    *Storefront
	MusicVideo    *MusicVideo
	LibrarySong   *LibrarySong
	LibraryAlbum  *LibraryAlbum
	Song          *Song
	Album         *Album
}

const CuratorType = "curators"
const AppleCuratorType = "apple-curators"
const ArtistType = "artists"
const LibraryArtistType = "library-artists"
const RatingType = "ratings"
const StationType = "station"
const StorefrontType = "storefronts"
const MusicVideoType = "music-videos"
const LibrarySongType = "library-songs"
const LibraryAlbumType = "library-albums"
const SongType = "songs"
const AlbumType = "albums"

func (r *SongOrMusicVideoResource) UnmarshalJSON(data []byte) error {
	aux := new(VariableResource)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Type = aux.Type
	r.Song = aux.Song
	r.MusicVideo = aux.MusicVideo
	return nil
}

func (r *VariableResource) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Type string
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	r.Type = aux.Type
	switch aux.Type {
	case CuratorType:
		var res = new(Curator)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.Curator = res
	case AppleCuratorType:
		var res = new(AppleCurator)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.AppleCurator = res
	case ArtistType:
		var res = new(Artist)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.Artist = res
	case LibraryArtistType:
		var res = new(LibraryArtist)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.LibraryArtist = res
	case RatingType:
		var res = new(Rating)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.Rating = res
	case StationType:
		var res = new(Station)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.Station = res
	case StorefrontType:
		var res = new(Storefront)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.Storefront = res
	case MusicVideoType:
		var res = new(MusicVideo)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.MusicVideo = res
	case LibrarySongType:
		var res = new(LibrarySong)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.LibrarySong = res
	case LibraryAlbumType:
		var res = new(LibraryAlbum)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.LibraryAlbum = res
	case SongType:
		var res = new(Song)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.Song = res
	case AlbumType:
		var res = new(Album)
		if err := json.Unmarshal(data, &res); err != nil {
			return err
		}
		r.Album = res
	default:
		return nil
	}

	return nil
}

type ReleaseDate struct {
	time.Time
}

// formatを設定
func (d ReleaseDate) format() string {
	return d.Time.Format("2006-01-02")
}

// MarshalJSON() の実装
func (self *ReleaseDate) UnmarshalJSON(data []byte) (err error) {
	s := string(data)

	// Get rid of the quotes "" around the value.
	// A second option would be to include them
	// in the date format string instead, like so below:
	//   time.Parse(`"`+time.RFC3339Nano+`"`, s)
	s = s[1 : len(s)-1]

	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t, err = time.Parse("2006-01-02", s)
	}
	self.Time = t
	return
}
