import React from 'react';
import {ButtonProps} from '../../types/Types';

export const Button: React.FC<ButtonProps> = ({content, onClick, className, variant = 'default'}) => {
	const style: string = `text-white text-sm px-3 py-1.5 rounded-md hover:opacity-70 transition hover:cursor-pointer font-semibold ${className} ${
		(variant === 'default' && 'bg-zinc-800') || (variant === 'success' && 'bg-green-500') || (variant === 'alert' && 'bg-red-800')
	}`;

	return (
		<button className={style} onClick={onClick}>
			{content}
		</button>
	);
};

export default Button;
