<template>
	<div class="w-100  mt-4">
		<div class="row px-3">
			<label class="col-1 col-form-label">Name</label>
			<label class="col-1 col-form-label">VO</label>
			<input v-model="feat.original_name" type="text" class="form-control col-4" placeholder="English Name">
			<label class="col-1 col-form-label">VF</label>
			<input v-model="feat.name" type="text" class="form-control col-4" placeholder="French Name">
			<div class="col-1 pr-0">
				<button v-if="feat.id" type="submit" class="form-control btn btn-success" @click="save(feat)">Update</button>
				<button v-if="!feat.id" type="submit" class="form-control btn btn-primary" @click="save(feat)">Add</button>
			</div>
		</div>

		<div class="row px-3 mt-2">
			<textarea v-model="feat.description" type="text" class="form-control col" rows="3" placeholder/>
		</div>
	</div>
</template>

<script>
export default {
	name: "Feat",
	props: ["id"],
	watch: {
		id: function(newVal) {
			this.getFeat(newVal);
		}
	},
	created: function() {
		this.getFeat(this.id);
	},
	data() {
		return {
			feat: {
				card_id: this.id
			},
			update: false
		};
	},
	methods: {
		getFeat: async function(cardID) {
			this.feat.card_id = cardID;
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/cards/" + cardID + "/feats")
				.then(function(res) {
					console.log(res);
					this.feat = res.data;
					if (!this.feat.id) {
						this.update = true;
					}
				});
		},
		save: function(feat) {
			if (feat.id == null) {
				feat.id = 0;
			}
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/feats/" + feat.id, feat)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						feat.id = res.data;
						this.update=false;
					}
				});
		}
	}
};
</script>

<style scoped>
</style>
