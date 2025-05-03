package main

import "gorm.io/gorm"

func FindPedagangAndLocation(gormdb *gorm.DB, args map[string]any) ([]PedagangWithJarak, error) {
	var pedagang []PedagangWithJarak
	query := gormdb.Raw(`
	SELECT
		pedagang.id,
		pedagang.nama,
		pedagang.jenis_dagangan,
		pedagang.foto_url,
		pedagang.jam_buka,
		pedagang.jam_tutup,
		pedagang.status,
		json_build_object(
			'id', lokasi.id,
			'latitude', lokasi.latitude,
			'longitude', lokasi.longitude,
			'alamat', lokasi.alamat,
			'kota', lokasi.kota
		) AS lokasi,

		-- Hitung jarak pakai rumus Haversine
		(6371 * acos(
			cos(radians(@latitude)) * cos(radians(lokasi.latitude)) * cos(radians(lokasi.longitude) - radians(@longitude)) +
			sin(radians(@latitude)) * sin(radians(lokasi.latitude))
		)) AS jarak_km

	FROM pedagang
	LEFT JOIN lokasi ON lokasi.pedagang_id = pedagang.id
	WHERE
	pedagang.nama ILIKE '%' || @q || '%'
	AND pedagang.jenis_dagangan ILIKE '%' || @jenis || '%'
	AND EXISTS (
		SELECT 1
		FROM pedagang_tag
		INNER JOIN tags ON tags.id = pedagang_tag.tag_id
		WHERE pedagang_tag.pedagang_id = pedagang.id
		AND tags.nama ILIKE '%' || @tag || '%'
	)
	AND lokasi.kota ILIKE '%' || @kota || '%'
	ORDER BY jarak_km ASC
	OFFSET @skip
	LIMIT @limit
	`, args)
	err := query.Scan(&pedagang).Error
	return pedagang, err
}
func FindPedagang(gormdb *gorm.DB, args map[string]any) ([]Pedagang, error) {
	var pedagang []Pedagang
	query := gormdb.Raw(`
	SELECT
		pedagang.id,
		pedagang.nama,
		pedagang.jenis_dagangan,
		pedagang.foto_url,
		pedagang.jam_buka,
		pedagang.jam_tutup,
		pedagang.status,
		json_build_object(
			'id', lokasi.id,
			'latitude', lokasi.latitude,
			'longitude', lokasi.longitude,
			'alamat', lokasi.alamat,
			'kota', lokasi.kota
		) AS lokasi
	FROM pedagang
	LEFT JOIN lokasi ON lokasi.pedagang_id = pedagang.id
	WHERE
	pedagang.nama ILIKE '%' || @q || '%'
	AND pedagang.jenis_dagangan ILIKE '%' || @jenis || '%'
	AND EXISTS (
		SELECT 1
		FROM pedagang_tag
		INNER JOIN tags ON tags.id = pedagang_tag.tag_id
		WHERE pedagang_tag.pedagang_id = pedagang.id
		AND tags.nama ILIKE '%' || @tag || '%'
	)
	AND lokasi.kota ILIKE '%' || @kota || '%'
	OFFSET @skip
	LIMIT @limit
	`, args)
	err := query.Scan(&pedagang).Error
	return pedagang, err
}
func FindPedagangById(gormdb *gorm.DB, id int) (PedagangDetail, error) {
	var pedagang PedagangDetail
	query := gormdb.Raw(`
	SELECT
	pedagang.id,
	pedagang.nama,
	pedagang.deskripsi,
	pedagang.jenis_dagangan,
	pedagang.foto_url,
	pedagang.jam_buka,
	pedagang.jam_tutup,
	pedagang.hari_mangkal,
	pedagang.status,
	pedagang.no_hp,
	pedagang.created_at,

	json_build_object(
		'id', lokasi.id,
		'pedagang_id', lokasi.pedagang_id,
		'latitude', lokasi.latitude,
		'longitude', lokasi.longitude,
		'alamat', lokasi.alamat,
		'kota', lokasi.kota
	) AS lokasi,

	COALESCE(
		json_agg(
		json_build_object(
			'id', tags.id,
			'nama', tags.nama
		)
		) FILTER (WHERE tags.id IS NOT NULL),
		'[]'::json
	) AS tags

	FROM pedagang
	LEFT JOIN lokasi ON lokasi.pedagang_id = pedagang.id
	LEFT JOIN pedagang_tag ON pedagang_tag.pedagang_id = pedagang.id
	LEFT JOIN tags ON tags.id = pedagang_tag.tag_id
	WHERE pedagang.id = ?
	GROUP BY pedagang.id, lokasi.id
	LIMIT 1;
 	`, id)
	err := query.Scan(&pedagang).Error
	return pedagang, err
}

func FindJenisDagangan(gormdb *gorm.DB, args map[string]any) ([]JenisDagangan, error) {
	var jenisDagangan []JenisDagangan
	query := gormdb.Raw(`
        SELECT DISTINCT jenis_dagangan AS value
        FROM pedagang
        WHERE jenis_dagangan ILIKE '%' || @q || '%'
        LIMIT @limit OFFSET @skip
	`, args).Scan(&jenisDagangan).Error
	return jenisDagangan, query
}
