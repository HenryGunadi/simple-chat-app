export const handleInput = <T,>(e: React.ChangeEvent<HTMLInputElement>, setState: React.Dispatch<React.SetStateAction<T>>) => {
	const value = e.target.value;
	setState(value as unknown as T);
};
