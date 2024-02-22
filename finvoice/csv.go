package finvoice

import (
	"encoding/csv"
	"io"
	"sync"
)

type CSVWriter struct {
	writer *csv.Writer

	headerOnce sync.Once
}

func NewCSVWriter(w io.Writer) *CSVWriter {
	return &CSVWriter{
		writer: csv.NewWriter(w),
	}
}

func (cw *CSVWriter) Write(inv *Finvoice) error {
	cw.headerOnce.Do(func() {
		cw.writeHeader()
	})

	for _, row := range inv.InvoiceRows {
		err := cw.writer.Write([]string{
			inv.InvoiceDetails.InvoiceNumber,
			inv.DeliveryDetails.DeliveryPeriodDetails.StartDate.Value,
			row.ArticleName,
			row.InvoicedQuantity.Value,
			row.UnitPriceAmount.Value,
			row.UnitPriceVatIncludedAmount.Value,
			row.RowDeliveryDate.Value,
			row.RowVatRatePercent,
			row.RowVatExcludedAmount.Value,
			row.RowAmount.Value,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (cw *CSVWriter) writeHeader() {
	header := []string{"InvoiceNumber", "InvoicePeriodStartDate", "ArticleName", "InvoicedQuantity", "UnitPriceAmount", "UnitPriceVatIncludedAmount", "RowDeliveryDate", "RowVatRatePercent", "RowVatExcludedAmount", "RowAmount"}
	cw.writer.Write(header)
}

func (cw *CSVWriter) Flush() {
	cw.writer.Flush()
}

func (cw *CSVWriter) Error() error {
	return cw.writer.Error()
}
