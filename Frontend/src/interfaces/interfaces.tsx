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

export interface FilterGroup {
  groupID: string;
  name: string;
  options: FilterOption[];
  changeSelectedOptions(
    selectedOptionsID: number[],
    changedOptionID: number
  ): void;
  returnType(): string;
}

export interface FilterOption {
  id: number;
  name: string;
}

export class CheckboxGroup implements FilterGroup {
  groupID: string;
  name: string;
  options: FilterOption[];

  constructor(groupID: string, name: string, options: FilterOption[]) {
    this.groupID = groupID;
    this.name = name;
    this.options = options;
  }
  returnType(): string {
    return "checkbox";
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

  constructor(groupID: string, name: string, options: FilterOption[]) {
    this.groupID = groupID;
    this.name = name;
    this.options = options;
  }
  returnType(): string {
    return "radio";
  }

  changeSelectedOptions(selectedOptionsID: number[], changedOptionID: number) {
    selectedOptionsID.splice(0, selectedOptionsID.length);
    selectedOptionsID.push(changedOptionID);
  }
}

export interface Exercise {
  id: number;
  name: string;
  muscles: Muscle[];
  difficulty: Difficulty;
  exerciseType: ExerciseType;
  equipment: Equipment[];
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
  type: string;
  difficulty: string;
  description: string;
  trainings: Training[];
}

export interface ProgramCard {
  id: number;
  name: string;
  type: string;
  difficulty: string;
  trainings: Training[];
}

export interface Training {
  id: number;
  name: string;
  trainingType: TrainingType;
  exercises: { exercise: Exercise; sets: string; reps: string; rest: string }[];
}
