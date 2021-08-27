import {PostModel} from './post.model';

export interface PostReactionsModel {
  username: string;
  likedPosts: PostModel[];
  dislikedPosts: PostModel[];
}


