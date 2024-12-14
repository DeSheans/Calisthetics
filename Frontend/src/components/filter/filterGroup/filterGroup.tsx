"use client";
import { FilterGroup as IFilterGroup } from "@/interfaces/interfaces";
import { Dispatch, SetStateAction, useState } from "react";
import styles from "./filterGroup.module.css";
import FilterOption from "../filterOption/filterOption";

const VIS_OPTIONS_DEFAULT = 5;
export default function FilterGroup({
  group,
  setLastChangedOption,
}: {
  group: IFilterGroup;
  setLastChangedOption: Dispatch<
    SetStateAction<
      | {
          groupID: string;
          optionID: number;
        }
      | undefined
    >
  >;
}) {
  const [visOptions, setVisOptions] = useState(() => VIS_OPTIONS_DEFAULT);

  return (
    <div className={styles.group}>
      <div className={styles.heading}>{group.name}</div>
      <div className={styles.optionContainer}>
        {group.options.slice(0, visOptions).map((option) => (
          <FilterOption
            key={option.id}
            option={option}
            group={group}
            setLastChangedOption={setLastChangedOption}
          ></FilterOption>
        ))}
      </div>

      {group.options.length > VIS_OPTIONS_DEFAULT &&
        GroupHider(setVisOptions, visOptions, group)}
    </div>
  );
}

function GroupHider(
  setVisOptions: Dispatch<SetStateAction<number>>,
  visOptions: number,
  group: IFilterGroup
) {
  return (
    <div className={styles.groupHider}>
      {
        <button
          className={styles.hiderButton}
          onClick={() =>
            setVisOptions(
              visOptions === VIS_OPTIONS_DEFAULT
                ? group.options.length
                : VIS_OPTIONS_DEFAULT
            )
          }
        >
          {visOptions > VIS_OPTIONS_DEFAULT ? "Скрыть" : "Показать всё"}
        </button>
      }
    </div>
  );
}
