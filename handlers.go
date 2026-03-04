package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func getAnsprechpartner(c *gin.Context) {
	limit, offset := getPaginationParams(c)
	query := `SELECT * FROM dbo.ANSPRECHPARTNER ORDER BY AdrNr ASC OFFSET ? ROWS FETCH NEXT ? ROWS ONLY`
	query = db.Rebind(query)

	var list []Ansprechpartner
	err := db.Select(&list, query, offset, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, list)
}

func getDomizil(c *gin.Context) {
	limit, offset := getPaginationParams(c)
	query := `SELECT * FROM dbo.DOMIZIL ORDER BY AdrNr ASC OFFSET ? ROWS FETCH NEXT ? ROWS ONLY`
	query = db.Rebind(query)

	var list []Domizil
	err := db.Select(&list, query, offset, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, list)
}

func getFall(c *gin.Context) {
	limit, offset := getPaginationParams(c)

	params := map[string]interface{}{
		"offset": offset,
		"limit":  limit,
	}

	query := "SELECT * FROM dbo.FALL WHERE 1=1"

	if owner := c.Query("owner"); owner != "" {
		query += " AND Owner = :owner"
		params["owner"] = owner
	}
	if status := c.Query("status"); status != "" {
		if status == "aktiv" {
			status = "1"
		} else if status == "inaktiv" {
			status = "2"
		}
		query += " AND Status = :status"
		params["status"] = status
	}
	if level := c.Query("level"); level != "" {
		query += " AND Fall_Level = :level"
		params["level"] = level
	}
	if dateFrom := c.Query("date_from"); dateFrom != "" {
		query += " AND FallBeginn >= :dateFrom"
		params["dateFrom"] = dateFrom
	}

	query += " ORDER BY FallNo ASC OFFSET :offset ROWS FETCH NEXT :limit ROWS ONLY"

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Error: " + err.Error()})
		return
	}

	query = db.Rebind(query)

	var list []Fall
	err = db.Select(&list, query, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": "Datenbankfehler: " + err.Error()})
		return
	}

	c.JSON(200, list)
}

func getSubjekt(c *gin.Context) {
	limit, offset := getPaginationParams(c)

	query := "SELECT * FROM dbo.SUBJEKT WHERE 1=1"
	params := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	if name := c.Query("name"); name != "" {
		query += " AND (Name1 LIKE :name OR Name2 LIKE :name)"
		params["name"] = "%" + name + "%"
	}
	if ort := c.Query("ort"); ort != "" {
		query += " AND Ort LIKE :ort"
		params["ort"] = "%" + ort + "%"
	}
	if plz := c.Query("plz"); plz != "" {
		query += " AND PLZ LIKE :plz"
		params["plz"] = plz + "%"
	}

	query += " ORDER BY AdrNr ASC OFFSET :offset ROWS FETCH NEXT :limit ROWS ONLY"

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Error: " + err.Error()})
		return
	}

	query = db.Rebind(query)

	var list []Subjekt
	err = db.Select(&list, query, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, list)
}

func getSubjektStruct(c *gin.Context) {
	limit, offset := getPaginationParams(c)
	query := `SELECT * FROM dbo.SUBJEKTSTRUCT ORDER BY AdrNr_1 ASC OFFSET ? ROWS FETCH NEXT ? ROWS ONLY`
	query = db.Rebind(query)

	var list []SubjektStruct
	err := db.Select(&list, query, offset, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, list)
}

func getFallDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID muss eine Zahl sein"})
		return
	}

	var f Fall
	query := db.Rebind("SELECT * FROM dbo.FALL WHERE FallNo = ?")
	err = db.Get(&f, query, id)

	if err != nil {
		c.JSON(404, gin.H{"error": "Fall nicht gefunden"})
		return
	}
	c.JSON(200, f)
}

func getSubjektDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID muss eine Zahl sein"})
		return
	}

	var s Subjekt
	query := db.Rebind("SELECT * FROM dbo.SUBJEKT WHERE AdrNr = ?")
	err = db.Get(&s, query, id)

	if err != nil {
		c.JSON(404, gin.H{"error": "Subjekt nicht gefunden"})
		return
	}
	c.JSON(200, s)
}

func getLog(c *gin.Context) {
	limit, offset := getPaginationParams(c)

	fallNo := c.Query("fall_no")
	logTyp := c.Query("type")
	owner := c.Query("owner")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	params := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	query := "SELECT * FROM dbo.LOG WHERE 1=1"

	if fallNo != "" {
		query += " AND Fall = :fallNo"
		params["fallNo"] = fallNo
	}
	if logTyp != "" {
		query += " AND LogTyp = :logTyp"
		params["logTyp"] = logTyp
	}
	if owner != "" {
		query += " AND Owner = :owner"
		params["owner"] = owner
	}
	if dateFrom != "" {
		query += " AND LogDatum >= :dateFrom"
		params["dateFrom"] = dateFrom
	}
	if dateTo != "" {
		query += " AND LogDatum <= :dateTo"
		params["dateTo"] = dateTo
	}

	query += " ORDER BY LogNo ASC OFFSET :offset ROWS FETCH NEXT :limit ROWS ONLY"

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Error: " + err.Error()})
		return
	}

	query = db.Rebind(query)

	var list []LogEintrag
	err = db.Select(&list, query, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": "Datenbankfehler: " + err.Error()})
		return
	}

	c.JSON(200, list)
}

func getLogDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID muss eine Zahl sein"})
		return
	}

	var l LogEintrag
	query := db.Rebind("SELECT * FROM dbo.LOG WHERE LogNo = ?")
	err = db.Get(&l, query, id)

	if err != nil {
		c.JSON(404, gin.H{"error": "Leistung (Log) nicht gefunden"})
		return
	}
	c.JSON(200, l)
}

func getPaginationParams(c *gin.Context) (int, int) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if err != nil || limit < 1 {
		limit = 100
	}
	if limit > 1000 {
		limit = 1000
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	return limit, offset
}
