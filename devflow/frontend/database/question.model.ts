type Question = {
  _id: string;
  title: string;
  description: string;
  userID: string;
  tags: string[];
  views: number;
  upvotes: string[];
  downvotes: string[];
  answers: string[]; //
  createdAt: string; //
};

export default Question;
