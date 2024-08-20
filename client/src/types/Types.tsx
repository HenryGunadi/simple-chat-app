import React from 'react';

export interface ButtonProps {
	content: string;
	onClick?: () => void;
	className?: string;
	variant?: 'default' | 'success' | 'alert';
}

export interface HttpRequestProps<T, D> {
	path: string;
	method: 'GET' | 'POST' | 'PATCH' | 'DELETE';
	data?: D;
	token: string;
	states: React.Dispatch<React.SetStateAction<T>>;
	navigate?: (path: string) => void;
}

export type LoginResponse = {
	status: string;
};

export type ChatType = {
	id: string;
	message: string;
};
