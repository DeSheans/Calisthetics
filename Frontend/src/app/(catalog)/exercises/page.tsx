"use client";
import Exercises from "@/components/exercises/exercises";
import Filter from "@/components/filter/filter";
import styles from "./page.module.css";
import { muscles } from "@/stubs/muscleStubs";
import { equipment } from "@/stubs/equipmentStub";
import { exerciseTypes } from "@/stubs/exerciseTypeStub";
import { difficulty } from "@/stubs/difficultyStub";
import {
  CheckboxGroup,
  FilterGroup,
  RadioGroup,
} from "@/interfaces/interfaces";
import { createContext, useState } from "react";

const filters = FetchFilterGroups();
const DataContext = createContext("");

export default function ExerciseCatalog() {
  const [filterParams, setFilterParams] = useState(() => "");

  return (
    <div className={styles.inner}>
      <h1 className={styles.heading}>Каталог упражнений</h1>
      <aside className={styles.filter}>
        <Filter
          FilterGroups={filters}
          setFilterParams={setFilterParams}
        ></Filter>
      </aside>
      <p>{filterParams}</p>
      <div className={styles.content}>
        <DataContext.Provider value={filterParams}>
          <Exercises></Exercises>
        </DataContext.Provider>
      </div>
    </div>
  );
}

function FetchFilterGroups(): FilterGroup[] {
  // fetch filters : {
  // filterGroupID = string
  // filterGroupName = string
  // filterOptions = {id:number, name:string}
  // }
  //

  return [
    new CheckboxGroup("muscles", "группы мышц", muscles),
    new CheckboxGroup("equipment", "оборудование", equipment),
    // new CheckboxGroup(2, "", [{ optionID:0, name: "избранное" }]),
    new RadioGroup("exerciseType", "тип упражнения", exerciseTypes),
    new RadioGroup("difficulty", "сложность", difficulty),
  ];
}
