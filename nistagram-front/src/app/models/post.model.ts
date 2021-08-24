export interface PostModel {
  numberOfComments: number;
  date: Date;
  description: string;
  NumberOfLikes: number;
  NumberOfDislikes: number;
  isAdd: boolean;
  isAlbum: boolean;
  NumberOfReaches: boolean;
  comments: object[];
  isPublic: boolean;
  location: string;
  tags: string[];
  path: string;
  owner: string;
  images: PostImages[];
}

interface PostImages{
  Image: string;
}
