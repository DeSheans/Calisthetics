import styles from "./filter.module.css";
import type { FilterGroup as IFilterGroup } from "@/interfaces/interfaces";
import FilterGroup from "./filterGroup/filterGroup";
import { Dispatch, SetStateAction, useEffect, useState } from "react";

export default function Filter({
  FilterGroups,
  setFilterParams,
}: {
  FilterGroups: IFilterGroup[];
  setFilterParams: Dispatch<SetStateAction<string>>;
}) {
  const [lastChangedOption, setLastChangedOption] = useState(
    (): { groupID: string; optionID: number } | undefined => undefined
  );

  useEffect(() => {
    if (lastChangedOption !== undefined) {
      setFilterParams(
        FilterParameterStringConstructor(
          FilterGroups,
          lastChangedOption.groupID,
          lastChangedOption.optionID
        )
      );
    }
  }, [lastChangedOption]);

  return (
    <div className={styles.filter}>
      <div className={styles.inner}>
        <div className={styles.heading}>
          <h2>Фильтр</h2>
        </div>

        {FilterGroups.map((g) => (
          <FilterGroup
            key={g.groupID}
            group={g}
            setLastChangedOption={setLastChangedOption}
          ></FilterGroup>
        ))}
      </div>
    </div>
  );
}

const filterParameterTable: {
  groupName: string;
  selectedOptionsID: number[];
}[] = [];

function FilterParameterStringConstructor(
  FilterList: IFilterGroup[],
  changedGroupName: string,
  changedOptionID: number
): string {
  let currentGroupFilter = filterParameterTable.find(
    (x) => x.groupName == changedGroupName
  );
  if (currentGroupFilter !== undefined) {
    let changedInputGroup = FilterList.find(
      (x) => x.groupID === changedGroupName
    );
    changedInputGroup?.changeSelectedOptions(
      currentGroupFilter.selectedOptionsID,
      changedOptionID
    );
  } else {
    filterParameterTable.push({
      groupName: changedGroupName,
      selectedOptionsID: [changedOptionID],
    });
  }

  let result: string[] = [];

  filterParameterTable.forEach((g) => {
    if (g.selectedOptionsID.length > 0) {
      result.push(`${g.groupName}=${g.selectedOptionsID.join(",")}`);
      result.push("&");
    }
  });
  result.pop();
  return "?" + result.join("");
}
