package persistent

const (
	insertQueryProduct = `
	INSERT INTO product(
		name, img
	) VALUES (
		$1,$2
	) RETURNING id;
	`
	updateQueryProduct = `
	UPDATE product SET 
		name = $1,
		img = $2
	WHERE id = $3
	`

	selectAllQueryProduct = `
		SELECT id, name, img
		FROM product;
	`

	selectByIdQueryProduct = `
		SELECT id, name, img
		FROM product
		WHERE id = $1;
	`
)
