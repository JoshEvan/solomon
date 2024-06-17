package persistent

const (
	insertQueryProduct = `
	INSERT INTO product(
		name, img, price
	) VALUES (
		$1,$2
	) RETURNING id;
	`

	updateQueryProduct = `
	UPDATE product SET 
		name = $1,
		img = $2,
		price = $3
	WHERE id = $4
	`

	selectAllQueryProduct = `
		SELECT id, name, img, price
		FROM product;
	`

	selectByIdQueryProduct = `
		SELECT id, name, img, price
		FROM product
		WHERE id = $1;
	`
)
