import { Exercise } from "@/interfaces/interfaces";
import {
  biceps,
  deltoids,
  lats,
  pectoral,
  quadriceps,
  rhomboid,
  trapezoidal,
  triceps,
} from "./muscleStubs";
import { pull, push, stat } from "./exerciseTypeStub";
import { easy, hard } from "./difficultyStub";

export const pushUp: Exercise = {
  id: 0,
  name: "Отжимания",
  muscles: [triceps, deltoids, pectoral],
  difficulty: easy,
  exerciseType: push,
  pictures: [],
  equipment: [],
  description: "",
  tips: [],
  variations: [],
};
export const pullUp: Exercise = {
  id: 1,
  name: "Подтягивания",
  muscles: [lats, trapezoidal, rhomboid],
  difficulty: easy,
  exerciseType: pull,
  pictures: [],
  equipment: [],
  description: "",
  tips: [],
  variations: [],
};
export const squat: Exercise = {
  id: 2,
  name: "Приседания",
  muscles: [quadriceps],
  difficulty: easy,
  exerciseType: push,
  pictures: [],
  equipment: [],
  description: "",
  tips: [],
  variations: [],
};
export const planche: Exercise = {
  id: 3,
  name: "Горизонт, планш",
  muscles: [biceps, triceps, deltoids],
  difficulty: hard,
  exerciseType: stat,
  pictures: [],
  equipment: [],
  description: "",
  tips: [],
  variations: [],
};
