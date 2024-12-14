import { ExerciseType } from "@/interfaces/interfaces";

export const pull = {
  id: 0,
  name: "тяга",
};

export const push = {
  id: 1,
  name: "жим",
};
export const stat = {
  id: 2,
  name: "статика",
};
export const cardio = {
  id: 3,
  name: "кардио",
};
export const flexibility = {
  id: 4,
  name: "растяжка",
};
export const balance = {
  id: 5,
  name: "баланс",
};

export const exerciseTypes: ExerciseType[] = [
  push,
  pull,
  stat,
  cardio,
  flexibility,
  balance,
];
