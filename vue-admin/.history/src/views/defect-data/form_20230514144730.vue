<template>
  <el-table
    :data="tableData"
    style="width: 100%"
    :default-sort="{ prop: 'date', order: 'descending' }"
  >
    <el-table-column prop="date" label="日期" sortable width="180" />
    <el-table-column prop="name" label="姓名" sortable width="180" />
    <el-table-column prop="address" label="地址" :formatter="formatter" />
  </el-table>
</template>
