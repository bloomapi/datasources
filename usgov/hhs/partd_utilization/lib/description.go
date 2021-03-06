package lib

import (
  "io"
  "regexp"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

var url = "http://download.cms.gov/Research-Statistics-Data-and-Systems/Statistics-Trends-and-Reports/Medicare-Provider-Charge-Data/Downloads/PartD_Prescriber_PUF_NPI_DRUG_13.zip"
var fileMatch = regexp.MustCompile(`PartD_Prescriber_PUF_NPI_DRUG_13.txt$`)

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.partd_utilization",
      Version: "2013-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  reader, err := getFileReader(url, fileMatch)
  if err != nil {
    return nil, err
  }

  csvReader := csv.NewReader(reader)
  if err != nil {
    return nil, err
  }

  csvReader.Comma = '\t'

  columns, err := csvReader.Read()
  if err != nil {
    return nil, err
  }

  return columns, nil
}

func getFileReader(uri string, zipPattern *regexp.Regexp) (io.Reader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
  if err != nil {
    return nil, err
  }

  reader, err := helpers.OpenExtractZipReader(path, zipPattern)
  if err != nil {
    return nil, err
  }

  return reader, nil
}

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
  reader, err := getFileReader(url, fileMatch)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvTabReader(reader)

  return zipReader, nil
}