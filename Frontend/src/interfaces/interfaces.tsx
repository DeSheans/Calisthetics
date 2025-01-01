export interface Muscle {
  id: number;
  name: string;
}

export interface Difficulty {
  id: number;
  name: string;
}

export interface ExerciseType {
  id: number;
  name: string;
}

export interface Equipment {
  id: number;
  name: string;
}

export interface TrainingType {
  id: number;
  type: string;
}

export interface ProgramType {
  id: number;
  name: string;
}

export interface StaticFilterGroup {
  groupID: string;
  name: string;
  options: FilterOption[];
  groupType: string;
}

export interface FilterGroup {
  groupID: string;
  name: string;
  options: FilterOption[];
  groupType: string;
  changeSelectedOptions(
    selectedOptionsID: number[],
    changedOptionID: number
  ): void;
}

export interface FilterOption {
  id: number;
  name: string;
}

export class CheckboxGroup implements FilterGroup {
  groupID: string;
  name: string;
  options: FilterOption[];
  groupType: string;

  constructor(
    groupID: string,
    name: string,
    options: FilterOption[],
    groupType: string
  ) {
    this.groupID = groupID;
    this.name = name;
    this.options = options;
    this.groupType = groupType;
  }

  changeSelectedOptions(selectedOptionsID: number[], changedOptionID: number) {
    let index = selectedOptionsID.findIndex((x) => x == changedOptionID);
    if (index !== -1) {
      selectedOptionsID.splice(index, 1);
    } else {
      selectedOptionsID.push(changedOptionID);
    }
  }
}

export class RadioGroup implements FilterGroup {
  groupID: string;
  name: string;
  options: FilterOption[];
  groupType: string;

  constructor(
    groupID: string,
    name: string,
    options: FilterOption[],
    groupType: string
  ) {
    this.groupID = groupID;
    this.name = name;
    this.options = options;
    this.groupType = groupType;
  }

  changeSelectedOptions(selectedOptionsID: number[], changedOptionID: number) {
    selectedOptionsID.splice(0, selectedOptionsID.length);
    selectedOptionsID.push(changedOptionID);
  }
}

export interface Exercise {
  id: number;
  name: string;
  muscles: string[];
  difficulty: string;
  exerciseType: string;
  equipment: string[];
  description: string;
  tips: string[];
  variations: string[];
  pictures: string[];
}

export interface ExerciseCard {
  id: number;
  name: string;
  muscles: string[];
  difficulty: string;
  exerciseType: string;
  picture: string;
}

export interface Program {
  id: number;
  name: string;
  programType: string;
  difficulty: string;
  description: string;
  trainings: Training[];
}

export interface ProgramCard {
  id: number;
  name: string;
  programType: string;
  difficulty: string;
  trainings: Training[];
}

export interface Training {
  id: number;
  name: string;
  trainingType: string;
  exercises: {
    exercise: string;
    sets: string;
    reps: string;
    rest: string;
  }[];
}
