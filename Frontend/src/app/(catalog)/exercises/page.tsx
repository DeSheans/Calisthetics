"use client";
import Exercises from "@/components/exercises/exercises";
import Filter from "@/components/filter/filter";
import styles from "./page.module.css";
import {
  CheckboxGroup,
  FilterGroup,
  RadioGroup,
  StaticFilterGroup,
} from "@/interfaces/interfaces";
import { createContext, useEffect, useState } from "react";

export const ExercisesContext = createContext("");

export default function ExerciseCatalog() {
  const [filterParams, setFilterParams] = useState(() => "");
  const [filters, setFilters] = useState((): FilterGroup[] => []);

  useEffect(() => {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    fetch(`http://${apiUrl}/exerciseFilter`)
      .then((res) => res.json())
      .then((data) => {
        setFilters(
          (data as StaticFilterGroup[]).map((g): FilterGroup => {
            return g.groupType === "checkbox"
              ? new CheckboxGroup(g.groupID, g.name, g.options, g.groupType)
              : new RadioGroup(g.groupID, g.name, g.options, g.groupType);
          })
        );
      });
  }, []);

  return (
    <div className={styles.inner}>
      <h1 className={styles.heading}>Каталог упражнений</h1>
      <aside className={styles.filter}>
        <Filter
          FilterGroups={filters}
          setFilterParams={setFilterParams}
        ></Filter>
      </aside>
      <div className={styles.content}>
        <ExercisesContext.Provider value={filterParams}>
          <Exercises></Exercises>
        </ExercisesContext.Provider>
      </div>
    </div>
  );
}
