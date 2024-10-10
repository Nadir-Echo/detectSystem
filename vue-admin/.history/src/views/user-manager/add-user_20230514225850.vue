<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="90px">
      <el-form-item label="用户名">
        <el-input v-model="form.username" />
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.password" autocomplete="off" />
      </el-form-item>
      <el-form-item label="姓名">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="负责生产线">
        <el-input v-model="form.proID" />
      </el-form-item>
      <el-form-item label="电话">
        <el-input v-model="form.phone" />
      </el-form-item>

      <el-form-item label="角色">
        <el-radio-group v-model="form.role">
          <el-radio label="admin" />
          <el-radio label="detector" />
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { addUser } from '@/api/manage'
export default {
  data() {
    return {
      form: {
        username: '',
        password: '',
        name: '',
        proID: null,
        phone: '',
        role: ''
      }
    }
  },
  methods: {
    onSubmit() {
      //   this.$message('submit!')
      console.log(this.form)
      addUser(this.form).then(res => {
        if (res.code === 200) {
          this.$message({
            message: '添加成功！',
            type: 'success'
          })
        } e 
        else {
          this.$message({
            message: '添加失败!',
            type: 'error'
          })
        }
      })
    },
    onCancel() {
      this.$message({
        message: '取消!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line {
  text-align: center;
}
</style>

