<template>
    <article class="px-2 mx-auto py-6 flex flex-col gap-6 border-b">
        <div class="space-y-6">
            <USkeleton class="h-6 w-full max-w-36" />

            <div class="space-y-4">
                <USkeleton class="h-4 w-full max-w-72" />
                <USkeleton class="h-4 w-full max-w-60" />
                <USkeleton class="h-4 w-full max-w-80" />
            </div>
        </div>

        <div class="flex justify-between items-center gap-4 mt-2">
            <div class="flex gap-2">
                <USkeleton class="h-10 w-10 rounded-full" />

                <div class="flex justify-center gap-1 flex-col">
                    <USkeleton class="h-4 w-[8rem]" />

                    <USkeleton class="h-4 w-[6rem]" />
                </div>
            </div>

            <nav class="flex gap-2">
                <USkeleton class="h-6 w-6" />

                <USkeleton class="h-6 w-6" />
            </nav>
        </div>
    </article>
</template>