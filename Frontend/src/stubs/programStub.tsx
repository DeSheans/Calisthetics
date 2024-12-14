import { Program, ProgramCard } from "@/interfaces/interfaces";
import { pull, push } from "./trainingTypeStub";
import { planche, pullUp, pushUp } from "./exerciseStubs";

export const twoDaySplit: Program = {
  id: 0,
  name: "Программа двух дневного сплита",
  type: "двух-дневный сплит",
  difficulty: "легко",
  description:
    "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Hic, magnam ratione recusandae rem illum similique enim. Repudiandae ratione quae nam eos, voluptate qui libero vitae mollitia, dignissimos aut iste. Magnam. Lorem ipsum dolor sit, amet consectetur adipisicing elit. Hic, magnam ratione recusandae rem illum similique enim. Repudiandae ratione quae nam eos, voluptate qui libero vitae mollitia, dignissimos aut iste. Magnam.Lorem ipsum dolor sit, amet consectetur adipisicing elit. Hic, magnam ratione recusandae rem illum similique enim. Repudiandae ratione quae nam eos, voluptate qui libero vitae mollitia, dignissimos aut iste. Magnam.",

  trainings: [
    {
      id: 0,
      name: "Тяговый день",
      trainingType: pull,
      exercises: [
        { exercise: pullUp, sets: "2", reps: "3-5", rest: "2 минуты" },
      ],
    },

    {
      id: 1,
      name: "День отжиманий",
      trainingType: push,
      exercises: [
        { exercise: planche, sets: "2", reps: "3-5", rest: "2 минуты" },
        { exercise: pushUp, sets: "3", reps: "5-7", rest: "1 минуты" },
      ],
    },
  ],
};

export const twoDaySplitCard: ProgramCard = {
  id: 0,
  name: "Программа двух дневного сплита",
  type: "двух-дневный сплит",
  difficulty: "легко",

  trainings: [
    {
      id: 0,
      name: "Тяговый день",
      trainingType: pull,
      exercises: [
        { exercise: pullUp, sets: "2", reps: "3-5", rest: "2 минуты" },
      ],
    },

    {
      id: 1,
      name: "День отжиманий",
      trainingType: push,
      exercises: [
        { exercise: planche, sets: "2", reps: "3-5", rest: "2 минуты" },
        { exercise: pushUp, sets: "3", reps: "5-7", rest: "1 минуты" },
      ],
    },
  ],
};
