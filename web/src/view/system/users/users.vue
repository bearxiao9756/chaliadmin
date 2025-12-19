
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

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
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
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
        
            <el-table-column align="left" label="用户ID" prop="id" width="120" />

            <el-table-column align="left" label="用户类型" prop="userType" width="120" />

            <el-table-column align="left" label="accessHash" prop="accessHash" width="120" />

            <el-table-column align="left" label="secretKey" prop="secretKeyId" width="120" />

            <el-table-column align="left" label="姓名-姓名" prop="firstName" width="120" />

            <el-table-column align="left" label="姓名-姓" prop="lastName" width="120" />

            <el-table-column align="left" label="姓名-氏" prop="username" width="120" />

            <el-table-column align="left" label="手机" prop="phone" width="120" />

            <el-table-column align="left" label="国家编码" prop="countryCode" width="120" />

            <el-table-column align="left" label="是否验证" prop="verified" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.verified) }}</template>
</el-table-column>
            <el-table-column align="left" label="是否支持" prop="support" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.support) }}</template>
</el-table-column>
            <el-table-column align="left" label="是否欺诈" prop="scam" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.scam) }}</template>
</el-table-column>
            <el-table-column align="left" label="是否假冒" prop="fake" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.fake) }}</template>
</el-table-column>
            <el-table-column align="left" label="是否特权" prop="premium" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.premium) }}</template>
</el-table-column>
            <el-table-column align="left" label="特权时间" prop="premiumExpireDate" width="120" />

            <el-table-column align="left" label="个人简介" prop="about" width="120" />

            <el-table-column align="left" label="用户状态" prop="state" width="120" />

            <el-table-column align="left" label="是否机器人" prop="isBot" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.isBot) }}</template>
</el-table-column>
            <el-table-column align="left" label="自动删除时间" prop="accountDaysTtl" width="120" />

            <el-table-column align="left" label="头像ID" prop="photoId" width="120" />

            <el-table-column align="left" label="是否封禁" prop="restricted" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.restricted) }}</template>
</el-table-column>
            <el-table-column align="left" label="封禁原因" prop="restrictionReason" width="120" />

            <el-table-column align="left" label="是否陌生人静音" prop="archiveAndMuteNewNoncontactPeers" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.archiveAndMuteNewNoncontactPeers) }}</template>
</el-table-column>
            <el-table-column align="left" label="表情文档ID" prop="emojiStatusDocumentId" width="120" />

            <el-table-column align="left" label="表情到期时间" prop="emojiStatusUntil" width="120" />

            <el-table-column align="left" label="最新的故事 ID" prop="storiesMaxId" width="120" />

            <el-table-column align="left" label="主颜色" prop="color" width="120" />

            <el-table-column align="left" label="表情背景色ID" prop="colorBackgroundEmojiId" width="120" />

            <el-table-column align="left" label="个人页面主色" prop="profileColor" width="120" />

            <el-table-column align="left" label="个人页面表情背景色ID" prop="profileColorBackgroundEmojiId" width="120" />

            <el-table-column align="left" label="生日" prop="birthday" width="120" />

            <el-table-column align="left" label="个人频道ID" prop="personalChannelId" width="120" />

            <el-table-column align="left" label="Token有效期" prop="authorizationTtlDays" width="120" />

            <el-table-column align="left" label="是否删除" prop="deleted" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.deleted) }}</template>
</el-table-column>
            <el-table-column align="left" label="删除原因" prop="deleteReason" width="120" />

            <el-table-column align="left" label="创建时间" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="更新时间" prop="updatedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.updatedAt) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateUsersFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
            <el-form-item label="用户ID:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入用户ID" />
</el-form-item>
            <el-form-item label="用户类型:" prop="userType">
    <el-input v-model.number="formData.userType" :clearable="true" placeholder="请输入用户类型" />
</el-form-item>
            <el-form-item label="accessHash:" prop="accessHash">
    <el-input v-model.number="formData.accessHash" :clearable="true" placeholder="请输入accessHash" />
</el-form-item>
            <el-form-item label="secretKey:" prop="secretKeyId">
    <el-input v-model.number="formData.secretKeyId" :clearable="true" placeholder="请输入secretKey" />
</el-form-item>
            <el-form-item label="姓名-姓名:" prop="firstName">
    <el-input v-model="formData.firstName" :clearable="true" placeholder="请输入姓名-姓名" />
</el-form-item>
            <el-form-item label="姓名-姓:" prop="lastName">
    <el-input v-model="formData.lastName" :clearable="true" placeholder="请输入姓名-姓" />
</el-form-item>
            <el-form-item label="姓名-氏:" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入姓名-氏" />
</el-form-item>
            <el-form-item label="手机:" prop="phone">
    <el-input v-model="formData.phone" :clearable="true" placeholder="请输入手机" />
</el-form-item>
            <el-form-item label="国家编码:" prop="countryCode">
    <el-input v-model="formData.countryCode" :clearable="true" placeholder="请输入国家编码" />
</el-form-item>
            <el-form-item label="是否验证:" prop="verified">
    <el-switch v-model="formData.verified" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="是否支持:" prop="support">
    <el-switch v-model="formData.support" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="是否欺诈:" prop="scam">
    <el-switch v-model="formData.scam" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="是否假冒:" prop="fake">
    <el-switch v-model="formData.fake" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="是否特权:" prop="premium">
    <el-switch v-model="formData.premium" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="特权时间:" prop="premiumExpireDate">
    <el-input v-model.number="formData.premiumExpireDate" :clearable="true" placeholder="请输入特权时间" />
</el-form-item>
            <el-form-item label="个人简介:" prop="about">
    <el-input v-model="formData.about" :clearable="true" placeholder="请输入个人简介" />
</el-form-item>
            <el-form-item label="用户状态:" prop="state">
    <el-input v-model.number="formData.state" :clearable="true" placeholder="请输入用户状态" />
</el-form-item>
            <el-form-item label="是否机器人:" prop="isBot">
    <el-switch v-model="formData.isBot" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="自动删除时间:" prop="accountDaysTtl">
    <el-input v-model.number="formData.accountDaysTtl" :clearable="true" placeholder="请输入自动删除时间" />
</el-form-item>
            <el-form-item label="头像ID:" prop="photoId">
    <el-input v-model.number="formData.photoId" :clearable="true" placeholder="请输入头像ID" />
</el-form-item>
            <el-form-item label="是否封禁:" prop="restricted">
    <el-switch v-model="formData.restricted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="封禁原因:" prop="restrictionReason">
    <el-input v-model="formData.restrictionReason" :clearable="true" placeholder="请输入封禁原因" />
</el-form-item>
            <el-form-item label="是否陌生人静音:" prop="archiveAndMuteNewNoncontactPeers">
    <el-switch v-model="formData.archiveAndMuteNewNoncontactPeers" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="表情文档ID:" prop="emojiStatusDocumentId">
    <el-input v-model.number="formData.emojiStatusDocumentId" :clearable="true" placeholder="请输入表情文档ID" />
</el-form-item>
            <el-form-item label="表情到期时间:" prop="emojiStatusUntil">
    <el-input v-model.number="formData.emojiStatusUntil" :clearable="true" placeholder="请输入表情到期时间" />
</el-form-item>
            <el-form-item label="最新的故事 ID:" prop="storiesMaxId">
    <el-input v-model.number="formData.storiesMaxId" :clearable="true" placeholder="请输入最新的故事 ID" />
</el-form-item>
            <el-form-item label="主颜色:" prop="color">
    <el-input v-model.number="formData.color" :clearable="true" placeholder="请输入主颜色" />
</el-form-item>
            <el-form-item label="表情背景色ID:" prop="colorBackgroundEmojiId">
    <el-input v-model.number="formData.colorBackgroundEmojiId" :clearable="true" placeholder="请输入表情背景色ID" />
</el-form-item>
            <el-form-item label="个人页面主色:" prop="profileColor">
    <el-input v-model.number="formData.profileColor" :clearable="true" placeholder="请输入个人页面主色" />
</el-form-item>
            <el-form-item label="个人页面表情背景色ID:" prop="profileColorBackgroundEmojiId">
    <el-input v-model.number="formData.profileColorBackgroundEmojiId" :clearable="true" placeholder="请输入个人页面表情背景色ID" />
</el-form-item>
            <el-form-item label="生日:" prop="birthday">
    <el-input v-model="formData.birthday" :clearable="true" placeholder="请输入生日" />
</el-form-item>
            <el-form-item label="个人频道ID:" prop="personalChannelId">
    <el-input v-model.number="formData.personalChannelId" :clearable="true" placeholder="请输入个人频道ID" />
</el-form-item>
            <el-form-item label="Token有效期:" prop="authorizationTtlDays">
    <el-input v-model.number="formData.authorizationTtlDays" :clearable="true" placeholder="请输入Token有效期" />
</el-form-item>
            <el-form-item label="是否删除:" prop="deleted">
    <el-switch v-model="formData.deleted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="删除原因:" prop="deleteReason">
    <el-input v-model="formData.deleteReason" :clearable="true" placeholder="请输入删除原因" />
</el-form-item>
            <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="更新时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="用户ID">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="用户类型">
    {{ detailForm.userType }}
</el-descriptions-item>
                    <el-descriptions-item label="accessHash">
    {{ detailForm.accessHash }}
</el-descriptions-item>
                    <el-descriptions-item label="secretKey">
    {{ detailForm.secretKeyId }}
</el-descriptions-item>
                    <el-descriptions-item label="姓名-姓名">
    {{ detailForm.firstName }}
</el-descriptions-item>
                    <el-descriptions-item label="姓名-姓">
    {{ detailForm.lastName }}
</el-descriptions-item>
                    <el-descriptions-item label="姓名-氏">
    {{ detailForm.username }}
</el-descriptions-item>
                    <el-descriptions-item label="手机">
    {{ detailForm.phone }}
</el-descriptions-item>
                    <el-descriptions-item label="国家编码">
    {{ detailForm.countryCode }}
</el-descriptions-item>
                    <el-descriptions-item label="是否验证">
    {{ detailForm.verified }}
</el-descriptions-item>
                    <el-descriptions-item label="是否支持">
    {{ detailForm.support }}
</el-descriptions-item>
                    <el-descriptions-item label="是否欺诈">
    {{ detailForm.scam }}
</el-descriptions-item>
                    <el-descriptions-item label="是否假冒">
    {{ detailForm.fake }}
</el-descriptions-item>
                    <el-descriptions-item label="是否特权">
    {{ detailForm.premium }}
</el-descriptions-item>
                    <el-descriptions-item label="特权时间">
    {{ detailForm.premiumExpireDate }}
</el-descriptions-item>
                    <el-descriptions-item label="个人简介">
    {{ detailForm.about }}
</el-descriptions-item>
                    <el-descriptions-item label="用户状态">
    {{ detailForm.state }}
</el-descriptions-item>
                    <el-descriptions-item label="是否机器人">
    {{ detailForm.isBot }}
</el-descriptions-item>
                    <el-descriptions-item label="自动删除时间">
    {{ detailForm.accountDaysTtl }}
</el-descriptions-item>
                    <el-descriptions-item label="头像ID">
    {{ detailForm.photoId }}
</el-descriptions-item>
                    <el-descriptions-item label="是否封禁">
    {{ detailForm.restricted }}
</el-descriptions-item>
                    <el-descriptions-item label="封禁原因">
    {{ detailForm.restrictionReason }}
</el-descriptions-item>
                    <el-descriptions-item label="是否陌生人静音">
    {{ detailForm.archiveAndMuteNewNoncontactPeers }}
</el-descriptions-item>
                    <el-descriptions-item label="表情文档ID">
    {{ detailForm.emojiStatusDocumentId }}
</el-descriptions-item>
                    <el-descriptions-item label="表情到期时间">
    {{ detailForm.emojiStatusUntil }}
</el-descriptions-item>
                    <el-descriptions-item label="最新的故事 ID">
    {{ detailForm.storiesMaxId }}
</el-descriptions-item>
                    <el-descriptions-item label="主颜色">
    {{ detailForm.color }}
</el-descriptions-item>
                    <el-descriptions-item label="表情背景色ID">
    {{ detailForm.colorBackgroundEmojiId }}
</el-descriptions-item>
                    <el-descriptions-item label="个人页面主色">
    {{ detailForm.profileColor }}
</el-descriptions-item>
                    <el-descriptions-item label="个人页面表情背景色ID">
    {{ detailForm.profileColorBackgroundEmojiId }}
</el-descriptions-item>
                    <el-descriptions-item label="生日">
    {{ detailForm.birthday }}
</el-descriptions-item>
                    <el-descriptions-item label="个人频道ID">
    {{ detailForm.personalChannelId }}
</el-descriptions-item>
                    <el-descriptions-item label="Token有效期">
    {{ detailForm.authorizationTtlDays }}
</el-descriptions-item>
                    <el-descriptions-item label="是否删除">
    {{ detailForm.deleted }}
</el-descriptions-item>
                    <el-descriptions-item label="删除原因">
    {{ detailForm.deleteReason }}
</el-descriptions-item>
                    <el-descriptions-item label="创建时间">
    {{ detailForm.createdAt }}
</el-descriptions-item>
                    <el-descriptions-item label="更新时间">
    {{ detailForm.updatedAt }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createUsers,
  deleteUsers,
  deleteUsersByIds,
  updateUsers,
  findUsers,
  getUsersList
} from '@/api/system/users'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'Users'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            id: undefined,
            userType: undefined,
            accessHash: undefined,
            secretKeyId: undefined,
            firstName: '',
            lastName: '',
            username: '',
            phone: '',
            countryCode: '',
            verified: false,
            support: false,
            scam: false,
            fake: false,
            premium: false,
            premiumExpireDate: undefined,
            about: '',
            state: undefined,
            isBot: false,
            accountDaysTtl: undefined,
            photoId: undefined,
            restricted: false,
            restrictionReason: '',
            archiveAndMuteNewNoncontactPeers: false,
            emojiStatusDocumentId: undefined,
            emojiStatusUntil: undefined,
            storiesMaxId: undefined,
            color: undefined,
            colorBackgroundEmojiId: undefined,
            profileColor: undefined,
            profileColorBackgroundEmojiId: undefined,
            birthday: '',
            personalChannelId: undefined,
            authorizationTtlDays: undefined,
            deleted: false,
            deleteReason: '',
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
    if (searchInfo.value.verified === ""){
        searchInfo.value.verified=null
    }
    if (searchInfo.value.support === ""){
        searchInfo.value.support=null
    }
    if (searchInfo.value.scam === ""){
        searchInfo.value.scam=null
    }
    if (searchInfo.value.fake === ""){
        searchInfo.value.fake=null
    }
    if (searchInfo.value.premium === ""){
        searchInfo.value.premium=null
    }
    if (searchInfo.value.isBot === ""){
        searchInfo.value.isBot=null
    }
    if (searchInfo.value.restricted === ""){
        searchInfo.value.restricted=null
    }
    if (searchInfo.value.archiveAndMuteNewNoncontactPeers === ""){
        searchInfo.value.archiveAndMuteNewNoncontactPeers=null
    }
    if (searchInfo.value.deleted === ""){
        searchInfo.value.deleted=null
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
  const table = await getUsersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteUsersFunc(row)
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
      const res = await deleteUsersByIds({ ids })
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
const updateUsersFunc = async(row) => {
    const res = await findUsers({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteUsersFunc = async (row) => {
    const res = await deleteUsers({ id: row.id })
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
        userType: undefined,
        accessHash: undefined,
        secretKeyId: undefined,
        firstName: '',
        lastName: '',
        username: '',
        phone: '',
        countryCode: '',
        verified: false,
        support: false,
        scam: false,
        fake: false,
        premium: false,
        premiumExpireDate: undefined,
        about: '',
        state: undefined,
        isBot: false,
        accountDaysTtl: undefined,
        photoId: undefined,
        restricted: false,
        restrictionReason: '',
        archiveAndMuteNewNoncontactPeers: false,
        emojiStatusDocumentId: undefined,
        emojiStatusUntil: undefined,
        storiesMaxId: undefined,
        color: undefined,
        colorBackgroundEmojiId: undefined,
        profileColor: undefined,
        profileColorBackgroundEmojiId: undefined,
        birthday: '',
        personalChannelId: undefined,
        authorizationTtlDays: undefined,
        deleted: false,
        deleteReason: '',
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
                  res = await createUsers(formData.value)
                  break
                case 'update':
                  res = await updateUsers(formData.value)
                  break
                default:
                  res = await createUsers(formData.value)
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
  const res = await findUsers({ id: row.id })
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
