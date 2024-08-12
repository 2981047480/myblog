<template>
    <div>
        <a-table :columns="columns" :data="data.Item" :loading="isLoading" show="true">

        </a-table>
    </div>
    <div>
    </div>
</template>

<script setup>
    import { ref, reactive, onMounted } from 'vue'
    import { GET_BLOG_LIST } from '../../../api/vblog/vblog';
    const columns = [
        {
            title: "id",
            dataIndex: "Id",
        },
        {
            title: "title",
            dataIndex: "title",
        },
        {
            title: "author",
            dataIndex: "author",
        },
        {
            title: "summary",
            dataIndex: "summary",
        },
        {
            title: "tags",
            dataIndex: "tags",
        },
    ];
    
    var isLoading = ref(false)
    var data = ref({ total: 0, item: [] })
    const Load = async () => {
        try {
            isLoading.value = true
            data = await GET_BLOG_LIST()
            console.log(data)
            console.log(data.Item,data.Total)
        }
        finally {
            isLoading.value = false
        }
    }

    onMounted(() => {
        Load()
    })

</script>

<style>
    
</style>