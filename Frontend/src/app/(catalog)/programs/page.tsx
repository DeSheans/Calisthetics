"use client";
import Filter from "@/components/filter/filter";
import Programs from "@/components/programs/programs";
import styles from "./page.module.css";
import {
  CheckboxGroup,
  FilterGroup,
  RadioGroup,
  StaticFilterGroup,
} from "@/interfaces/interfaces";
import { createContext, useEffect, useState } from "react";

export const ProgramsContext = createContext("");

export default function ProgramsCatalog() {
  const [filterParams, setFilterParams] = useState(() => "");
  const [filters, setFilters] = useState((): FilterGroup[] => []);

  useEffect(() => {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    fetch(`http://${apiUrl}/programFilter`)
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
      <h1 className={styles.heading}>Каталог программ</h1>
      <div className={styles.filter}>
        <Filter
          FilterGroups={filters}
          setFilterParams={setFilterParams}
        ></Filter>
      </div>
      <div className={styles.content}>
        <ProgramsContext.Provider value={filterParams}>
          <Programs></Programs>
        </ProgramsContext.Provider>
      </div>
    </div>
  );
}
