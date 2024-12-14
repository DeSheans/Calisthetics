import { ExerciseCard } from "@/interfaces/interfaces";

export const pushUps: ExerciseCard = {
  id: 0,
  name: "Отжимания",
  muscles: ["трицепс", "дельтовидные", "грудные"],
  difficulty: "легко",
  exerciseType: "жим",
  picture: "src\\assets\\pictures\\",
};
export const pullUps: ExerciseCard = {
  id: 1,
  name: "Подтягивания",
  muscles: ["широчайшие", "трапециевидные", "ромбовидные"],
  difficulty: "легко",
  exerciseType: "тяга",
  picture: "src\\assets\\pictures\\",
};
export const squats: ExerciseCard = {
  id: 2,
  name: "Приседания",
  muscles: ["квадрицепс"],
  difficulty: "легко",
  exerciseType: "жим",
  picture: "src\\assets\\pictures\\",
};
export const planche: ExerciseCard = {
  id: 3,
  name: "Горизонт, планш",
  muscles: ["бицепс", "трицепс", "дельтовидные"],
  difficulty: "сложно",
  exerciseType: "статика",
  picture: "src\\assets\\pictures\\",
};

export const exercises = [pushUps, pullUps, squats, planche];
