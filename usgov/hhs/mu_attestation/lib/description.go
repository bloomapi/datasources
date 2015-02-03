package lib

import (
  "os"
  "encoding/csv"
  "github.com/gocodo/bloomsource"
  "github.com/gocodo/bloomsource/helpers"
)

type Description struct {}

func (d *Description) Available() ([]bloomsource.Source, error) {
  return []bloomsource.Source{
    bloomsource.Source{
      Name: "usgov.hhs.mu_attestation",
      Version: "2015-01-20",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch("http://s3.amazonaws.com/gocodo/usgov/hhs/20150120+EH+Attestation+Summary+-+Final+-+Stage+1.csv")
  if err != nil {
    return nil, err
  }

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer fileReader.Close()

  csvReader := csv.NewReader(fileReader)
  if err != nil {
    return nil, err
  }

  columns, err := csvReader.Read()
  if err != nil {
    return nil, err
  }

  return columns, nil
}

func (d *Description) Reader(source bloomsource.Source) (bloomsource.ValueReader, error) {
  downloader := bloomsource.NewDownloader("data/", nil)
  path, err := downloader.Fetch("http://s3.amazonaws.com/gocodo/usgov/hhs/20150120+EH+Attestation+Summary+-+Final+-+Stage+1.csv")
  if err != nil {
    return nil, err
  }

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  csvReader := helpers.NewCsvReader(fileReader)

  return csvReader, nil
}