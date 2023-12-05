type User = {
  _id: string;
  clerkID: string;
  firstName: string;
  lastName: string;
  bio: string;
  picture: string;
  email: string;
  location: string;
  portfolioWebsite: string;
  isAdmin: boolean;
  reputation: number;
  saved: string[];
  joinedAt: string;
};

export default User;