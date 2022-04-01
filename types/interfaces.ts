export interface House {
	id: number;
	title: string;
	type: 'house' | 'apartment';
	location: string;
	price: number;
	imageUrl: string;
	owner: string;
}
