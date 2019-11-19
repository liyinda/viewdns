<template>
  <div class="app-container">
    <el-form ref="dataForm" :model="form" label-width="120px">
      <el-form-item label="Domain Name">
        <el-input v-model="form.domain" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">确认删除</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { delDomain } from '@/api/table'

export default {
  data() {
    return {
      form: {
        domain: ''
      }
    }
  },
  methods: {
    onSubmit() {
      this.$refs.dataForm.validate((valid) => {
        if (valid) {
          delDomain(this.form).then(() => {
            this.dialogFormVisible = false
            this.$message({
              message: '删除成功',
              type: 'warning'
            })
          })
        }
      })
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

