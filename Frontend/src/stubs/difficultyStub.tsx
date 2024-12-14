import { Difficulty } from "@/interfaces/interfaces";

export const easy: Difficulty = {
  id: 0,
  name: "легко",
};

export const normal: Difficulty = {
  id: 1,
  name: "средне",
};

export const hard: Difficulty = {
  id: 2,
  name: "сложно",
};

export const difficulty: Difficulty[] = [easy, normal, hard];
