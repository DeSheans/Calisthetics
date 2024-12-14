"use client";
import Filter from "@/components/filter/filter";
import Programs from "@/components/programs/programs";
import styles from "./page.module.css";
import {
  CheckboxGroup,
  FilterGroup,
  RadioGroup,
} from "@/interfaces/interfaces";
import { difficulty } from "@/stubs/difficultyStub";
import { programTypes } from "@/stubs/programTypeStub";
import { createContext, useState } from "react";

const filters = FetchFilterGroups();
const DataContext = createContext("");

export default function ProgramsCatalog() {
  const [filterParams, setFilterParams] = useState(() => "");

  return (
    <div className={styles.inner}>
      <h1 className={styles.heading}>Каталог программ</h1>
      <div className={styles.filter}>
        <Filter
          FilterGroups={filters}
          setFilterParams={setFilterParams}
        ></Filter>
      </div>
      <div className={styles.content}>
        <DataContext.Provider value={filterParams}>
          <Programs></Programs>
        </DataContext.Provider>
      </div>
    </div>
  );
}

function FetchFilterGroups(): FilterGroup[] {
  return [
    new RadioGroup("programTypes", "тип программы", programTypes),
    // new CheckboxGroup("favorite", "", [{ id: 0, name: "избранное" }]),
    new RadioGroup("difficulty", "сложность", difficulty),
  ];
}
