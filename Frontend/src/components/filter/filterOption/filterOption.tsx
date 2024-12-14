import {
  FilterGroup,
  FilterOption as IFilterOption,
} from "@/interfaces/interfaces";
import styles from "./filterOption.module.css";
import { Dispatch, SetStateAction } from "react";
export default function FilterOption({
  option,
  group,
  setLastChangedOption,
}: {
  option: IFilterOption;
  group: FilterGroup;
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
  return (
    <label className={styles.option}>
      <input
        className={styles.optionInput}
        type={group.returnType()}
        name={group.groupID}
        id={option.id.toString()}
        onChange={() =>
          setLastChangedOption({
            groupID: group.groupID,
            optionID: option.id,
          })
        }
      ></input>
      <span className={styles.optionSpan}>{option.name}</span>
    </label>
  );
}
