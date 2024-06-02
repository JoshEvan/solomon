package persistent

const (
	upsertQueryProduct = `
	INSERT INTO product(
		id, name, img,
		metadata
	) VALUES (
		$1,$2,$3,$4
	) ON CONFLICT(id)
	DO UPDATE SET
		name = EXLCUDED.name,
		price = EXCLUDED.price,
		metadata = EXCLUDED.metadata;
	`

	selectAllQueryProduct = `
		SELECT 
			id, name, img, metadata
		FROM product;
	`
)
