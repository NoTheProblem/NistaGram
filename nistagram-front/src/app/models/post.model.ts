export interface PostModel {
  id: string;
  numberOfComments: number;
  date: Date;
  description: string;
  NumberOfLikes: number;
  NumberOfDislikes: number;
  usersLiked: string[];
  usersDisliked: string[];
  isAdd: boolean;
  isAlbum: boolean;
  NumberOfReaches: boolean;
  comments: Comment[];
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


interface Comment{
  text: string;
  date: Date;
  commentOwnerUsername: string;
}
