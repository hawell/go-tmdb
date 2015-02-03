package tmdb

import (
	. "gopkg.in/check.v1"
)

const gameOfThronesID int = 1399

func (s *TmdbSuite) TestGetTvInfo(c *C) {
	result, err := s.tmdb.GetTvInfo(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Name, Equals, "Game of Thrones")
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.FirstAirDate, Equals, "2011-04-17")
	c.Assert(result.Homepage, Equals, "http://www.hbo.com/game-of-thrones/index.html")
	c.Assert(result.Networks[0].ID, Equals, hboID)
	c.Assert(result.Networks[0].Name, Equals, "HBO")
	c.Assert(result.Status, Equals, "Returning Series")
	c.Assert(result.Type, Equals, "Scripted")
	c.Assert(len(result.Seasons), Equals, result.NumberOfSeasons+1)

	var options = make(map[string]string)
	// options["append_to_response"] = "account_states,alternative_titles,credits,images,keywords,releases,videos,translations,similar,reviews,lists,changes,ratings"
	options["append_to_response"] = "account_states,alternative_titles,changes"
	allResult, err := s.tmdb.GetTvInfo(gameOfThronesID, options)
	s.baseTest(&allResult, err, c)
	c.Assert(allResult.ID, Equals, gameOfThronesID)
	c.Assert(allResult.Name, Equals, "Game of Thrones")
	c.Assert(allResult.AlternativeTitles, NotNil)
	c.Assert(allResult.Changes, NotNil)
	// c.Assert(allResult.Credits, NotNil)
	// c.Assert(allResult.Images, NotNil)
	// c.Assert(allResult.Keywords, NotNil)
	// c.Assert(allResult.Releases, NotNil)
	// c.Assert(allResult.Videos, NotNil)
	// c.Assert(allResult.Translations, NotNil)
	// c.Assert(allResult.Similar, NotNil)
	// c.Assert(allResult.Reviews, NotNil)
	// c.Assert(allResult.Lists, NotNil)
}

func (s *TmdbSuite) TestGetTvAccountStates(c *C) {
	// result, err := s.tmdb.GetTvAccountStates(gameOfThronesID)
	// s.baseTest(&result, err, c)
	// TODO
}

func (s *TmdbSuite) TestGetTvAlternativeTitles(c *C) {
	result, err := s.tmdb.GetTvAlternativeTitles(gameOfThronesID)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvChanges(c *C) {
	result, err := s.tmdb.GetTvChanges(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Changes, NotNil)
	allResultLength := len(result.Changes)

	var janOptions = make(map[string]string)
	janOptions["start_date"] = "2015-01-01"
	janOptions["end_date"] = "2015-01-14"
	janResult, err := s.tmdb.GetTvChanges(gameOfThronesID, janOptions)
	s.baseTest(&janResult, err, c)
	c.Assert(janResult.Changes, NotNil)
	c.Assert(len(janResult.Changes) >= allResultLength, Equals, true)
}
