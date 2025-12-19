
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
            <el-form-item label="编号" prop="id">
  <el-input v-model.number="searchInfo.id" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="机器人ID" prop="botId">
  <el-input v-model.number="searchInfo.botId" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="机器人类型" prop="botType">
  <el-input v-model.number="searchInfo.botType" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="创作者用户 ID" prop="creatorUserId">
  <el-input v-model.number="searchInfo.creatorUserId" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="机器人聊天记录" prop="botChatHistory">
  <el-select v-model="searchInfo.botChatHistory" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
            
            <el-form-item label="机器人无聊天" prop="botNochats">
  <el-select v-model="searchInfo.botNochats" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
            
            <el-form-item label="是否验证" prop="verified">
  <el-select v-model="searchInfo.verified" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
            
            <el-form-item label="botInlineGeo字段" prop="botInlineGeo">
  <el-select v-model="searchInfo.botInlineGeo" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
</el-form-item>
            
            <el-form-item label="botInfoVersion字段" prop="botInfoVersion">
  <el-input v-model.number="searchInfo.botInfoVersion" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="创建于" prop="createdAt">
<el-date-picker v-model="searchInfo.createdAt" type="datetime" placeholder="搜索条件"></el-date-picker></el-form-item>
            
            <el-form-item label="创建于" prop="updatedAt">
<el-date-picker v-model="searchInfo.updatedAt" type="datetime" placeholder="搜索条件"></el-date-picker></el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="编号" prop="id" width="120" />

            <el-table-column align="left" label="机器人ID" prop="botId" width="120" />

            <el-table-column align="left" label="机器人类型" prop="botType" width="120" />

            <el-table-column align="left" label="创作者用户 ID" prop="creatorUserId" width="120" />

            <el-table-column align="left" label="令牌" prop="token" width="120" />

            <el-table-column align="left" label="描述" prop="description" width="120" />

            <el-table-column align="left" label="机器人聊天记录" prop="botChatHistory" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.botChatHistory) }}</template>
</el-table-column>
            <el-table-column align="left" label="机器人无聊天" prop="botNochats" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.botNochats) }}</template>
</el-table-column>
            <el-table-column align="left" label="是否验证" prop="verified" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.verified) }}</template>
</el-table-column>
            <el-table-column align="left" label="botInlineGeo字段" prop="botInlineGeo" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.botInlineGeo) }}</template>
</el-table-column>
            <el-table-column align="left" label="botInfoVersion字段" prop="botInfoVersion" width="120" />

            <el-table-column align="left" label="botInlinePlaceholder字段" prop="botInlinePlaceholder" width="120" />

            <el-table-column align="left" label="创建于" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="创建于" prop="updatedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateTeamChaliBotsFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="编号:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入编号" />
</el-form-item>
            <el-form-item label="机器人ID:" prop="botId">
    <el-input v-model.number="formData.botId" :clearable="true" placeholder="请输入机器人ID" />
</el-form-item>
            <el-form-item label="机器人类型:" prop="botType">
    <el-input v-model.number="formData.botType" :clearable="true" placeholder="请输入机器人类型" />
</el-form-item>
            <el-form-item label="创作者用户 ID:" prop="creatorUserId">
    <el-input v-model.number="formData.creatorUserId" :clearable="true" placeholder="请输入创作者用户 ID" />
</el-form-item>
            <el-form-item label="令牌:" prop="token">
    <el-input v-model="formData.token" :clearable="true" placeholder="请输入令牌" />
</el-form-item>
            <el-form-item label="描述:" prop="description">
    <el-input v-model="formData.description" :clearable="true" placeholder="请输入描述" />
</el-form-item>
            <el-form-item label="机器人聊天记录:" prop="botChatHistory">
    <el-switch v-model="formData.botChatHistory" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="机器人无聊天:" prop="botNochats">
    <el-switch v-model="formData.botNochats" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="是否验证:" prop="verified">
    <el-switch v-model="formData.verified" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="botInlineGeo字段:" prop="botInlineGeo">
    <el-switch v-model="formData.botInlineGeo" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="botInfoVersion字段:" prop="botInfoVersion">
    <el-input v-model.number="formData.botInfoVersion" :clearable="true" placeholder="请输入botInfoVersion字段" />
</el-form-item>
            <el-form-item label="botInlinePlaceholder字段:" prop="botInlinePlaceholder">
    <el-input v-model="formData.botInlinePlaceholder" :clearable="true" placeholder="请输入botInlinePlaceholder字段" />
</el-form-item>
            <el-form-item label="创建于:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="创建于:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="编号">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="机器人ID">
    {{ detailForm.botId }}
</el-descriptions-item>
                    <el-descriptions-item label="机器人类型">
    {{ detailForm.botType }}
</el-descriptions-item>
                    <el-descriptions-item label="创作者用户 ID">
    {{ detailForm.creatorUserId }}
</el-descriptions-item>
                    <el-descriptions-item label="令牌">
    {{ detailForm.token }}
</el-descriptions-item>
                    <el-descriptions-item label="描述">
    {{ detailForm.description }}
</el-descriptions-item>
                    <el-descriptions-item label="机器人聊天记录">
    {{ detailForm.botChatHistory }}
</el-descriptions-item>
                    <el-descriptions-item label="机器人无聊天">
    {{ detailForm.botNochats }}
</el-descriptions-item>
                    <el-descriptions-item label="是否验证">
    {{ detailForm.verified }}
</el-descriptions-item>
                    <el-descriptions-item label="botInlineGeo字段">
    {{ detailForm.botInlineGeo }}
</el-descriptions-item>
                    <el-descriptions-item label="botInfoVersion字段">
    {{ detailForm.botInfoVersion }}
</el-descriptions-item>
                    <el-descriptions-item label="botInlinePlaceholder字段">
    {{ detailForm.botInlinePlaceholder }}
</el-descriptions-item>
                    <el-descriptions-item label="创建于">
    {{ detailForm.createdAt }}
</el-descriptions-item>
                    <el-descriptions-item label="创建于">
    {{ detailForm.updatedAt }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createTeamChaliBots,
  deleteTeamChaliBots,
  deleteTeamChaliBotsByIds,
  updateTeamChaliBots,
  findTeamChaliBots,
  getTeamChaliBotsList
} from '@/api/system/teamchalibots'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'TeamChaliBots'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            botId: undefined,
            botType: undefined,
            creatorUserId: undefined,
            token: '',
            description: '',
            botChatHistory: false,
            botNochats: false,
            verified: false,
            botInlineGeo: false,
            botInfoVersion: undefined,
            botInlinePlaceholder: '',
            createdAt: new Date(),
            updatedAt: new Date(),
        })



// 验证规则
const rule = reactive({
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.botChatHistory === ""){
        searchInfo.value.botChatHistory=null
    }
    if (searchInfo.value.botNochats === ""){
        searchInfo.value.botNochats=null
    }
    if (searchInfo.value.verified === ""){
        searchInfo.value.verified=null
    }
    if (searchInfo.value.botInlineGeo === ""){
        searchInfo.value.botInlineGeo=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getTeamChaliBotsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTeamChaliBotsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteTeamChaliBotsByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTeamChaliBotsFunc = async(row) => {
    const res = await findTeamChaliBots({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTeamChaliBotsFunc = async (row) => {
    const res = await deleteTeamChaliBots({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        id: undefined,
        botId: undefined,
        botType: undefined,
        creatorUserId: undefined,
        token: '',
        description: '',
        botChatHistory: false,
        botNochats: false,
        verified: false,
        botInlineGeo: false,
        botInfoVersion: undefined,
        botInlinePlaceholder: '',
        createdAt: new Date(),
        updatedAt: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createTeamChaliBots(formData.value)
                  break
                case 'update':
                  res = await updateTeamChaliBots(formData.value)
                  break
                default:
                  res = await createTeamChaliBots(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findTeamChaliBots({ id: row.id })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
