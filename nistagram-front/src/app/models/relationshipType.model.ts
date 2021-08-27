export interface RelationType {
  relation: RELATION_TYPE;
}

export enum RELATION_TYPE  {
  NOT_FOLLOWING,
  FOLLOWING ,
  NOT_ACCEPTED ,
  BLOCKED,
  BLOCKING
}

