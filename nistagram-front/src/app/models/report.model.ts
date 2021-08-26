import {PostModel} from './post.model';

export interface ReportModel {
  id: string;
  postId: string;
  penalty: PENALTY;
  date: Date;
  isAnswered: boolean;
  post: PostModel;
}

export enum PENALTY  {
  REMOVE_CONTENT,
  DELETE_PROFILE,
  DECLINE_REPORT
}
