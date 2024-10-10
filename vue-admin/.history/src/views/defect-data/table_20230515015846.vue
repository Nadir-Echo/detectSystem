<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="PID" width="95">
        <template slot-scope="scope">
          {{ scope.row.pid }}
        </template>
      </el-table-column>
      <el-table-column label="缺陷" align="center">
        <template slot-scope="scope">
          {{ scope.row.class }}
        </template>
      </el-table-column>
      <el-table-column label="生产线" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.proId }}</span>
        </template>
      </el-table-column>

      <el-table-column
        align="center"
        prop="created_at"
        label="Defect_time"
        width="200"
      >
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.time }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="Detail">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleDetail(scope.row)"
            >查看详情</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <!-- 详情弹窗 -->
    <el-dialog :visible.sync="dialogVisible" title="详情" width="50%">
      <el-form :model="detailForm" label-width="80px" class="detail-form">
        <el-form-item label="PID">
          <span>{{ detailForm.pid }}</span>
        </el-form-item>
        <el-form-item label="缺陷">
          <span>{{ detailForm.class }}</span>
        </el-form-item>
        <el-form-item label="生产线">
          <span>{{ detailForm.proId }}</span>
        </el-form-item>
        <el-form-item label="Defect_time">
          <span>{{ detailForm.time }}</span>
        </el-form-item>
        <!-- 添加其他需要显示的字段 -->
        <el-row>
          <el-col :span="12">
            <el-form-item label="图片1">
              <el-image :src=src1 fit="contain" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="字段1">
              <span>{{ detailForm.field1 }}</span>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="图片2">
              <el-image :src=src2 fit="contain" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="字段2">
              <span>{{ detailForm.field2 }}</span>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>

import { getDataForm } from '@/api/defect-data'
export default {
  data() {
    return {
      list: null,
      listLoading: true,
      dialogVisible: false, // 控制详情弹窗的显示
      detailForm: {} // 用于存储详情数据的对象
      src1
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getDataForm().then(response => {
        this.list = response.data
        this.listLoading = false
      })
    },
    handleDetail(row) {
      this.$alert(`PID: ${row.pid}, 缺陷: ${row.class}, 生产线: ${row.proId}, Defect_time: ${row.time},src1:"https://joplin-1305917469.cos.ap-nanjing.myqcloud.com/Snipaste_2023-04-17_16-23-26.png"`, '详情', {
        confirmButtonText: '确定',
        dangerouslyUseHTMLString: true
      })
    }
  }
}
</script>
