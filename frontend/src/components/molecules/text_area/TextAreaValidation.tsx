import { FieldAttributes, useField } from "formik";
import React from 'react';
import { TextareaAutosize } from "@material-ui/core";

export const TextAreaWValidation:React.FC<FieldAttributes<{}>> = ({placeholder,...props}) => {
	const [field, meta] = useField<{}>(props); // hook dari formik
	const errorText = meta.error && meta.touched ? meta.error : "";
	return(
		<TextareaAutosize aria-label={placeholder} rowsMin={3} placeholder={placeholder} {...field} 
		helperText={errorText} error={!!errorText} />
	)
}