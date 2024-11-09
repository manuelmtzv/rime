export type User = {
  id: string;
  name: string;
  lastname: string;
  username: string;
  email: string;
  password: string;
  createdAt: string;
  updatedAt: string;
};

export type PopularUser = Omit<User, "createdAt" | "updatedAt"> & {
  followers: number;
};
