<template>
	<div class="w-100">
		<div v-if="!update" class="row px-3">
			<span class="col-2 text-left">{{selectedAbility.name}} ({{selectedAbility.original_name}})</span>
			<span class="col-9 text-left">{{selectedAbility.description}}</span>
			<span class="col-1 form-inline text-right">
				<button type="submit" class="btn-sm btn-success" @click="update = true">U</button>
				<button type="submit" class="btn-sm btn-danger" @click="$emit('remove')">X</button>
			</span>
		</div>

		<div v-if="update" class="row px-3">
			<v-autocomplete
				v-if="!ability.id"
				:items="items"
				:get-label="getLabel"
				@update-items="updateItems"
				:component-item="template"
				:auto-select-one-item="false"
				@item-selected="selectedItem"
				class="col-2 mt-1"
			></v-autocomplete>
			<input v-model="selectedAbility.original_name" type="text" class="form-control col-2" placeholder="English Name">
			<input v-model="selectedAbility.name" type="text" class="form-control col-2" placeholder="French Name">
			<div class="form-check form-check-inline ml-2 col-5">
				<label class="form-check-label">Can be magical or choose type ?</label>
				<input class="form-check-input ml-3" type="checkbox" v-model="selectedAbility.magical">
			</div>
				<textarea v-model="selectedAbility.description" type="text" class="form-control col-11" rows="3" placeholder/>
				<div class="col-1 px-0">
					<button v-if="ability.id || selectedAbility.id" type="submit" class="form-control btn btn-success" @click="save(selectedAbility)">Update</button>
					<button v-if="ability.id" type="submit" class="form-control btn btn-danger" @click="remove(selectedAbility)">Delete</button>
					<button v-if="!ability.id && selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="add(selectedAbility)">Add</button>
					<button v-if="!ability.id && !selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="save(selectedAbility)">Add</button>
				</div>
		</div>
	</div>
</template>

<script>
import ItemTemplate from "./ItemTemplate.vue";
export default {
	name: "Ability",
	props: ["abilitiesList", "ability"],
	watch: {
		ability: function(newVal, oldVal) {
			console.log("CHANGED ABILITY", newVal, oldVal);
			this.selectedAbility = newVal;
			if (!this.ability.id){
				this.update=true;
			}
		}
	},
	created: function() {
		this.selectedAbility = this.ability;
		console.log("ON CERATE ABILITY ID", this.ability)
		if (!this.ability.id){
				this.update=true;
			}

	},
	data() {
		return {
			selectedAbility: {},
			template: ItemTemplate,
			items: [],
			update: false
		};
	},
	methods: {
		save: function(ability) {
			if (ability.id == null) {
				ability.id = 0;
			}
			this.$http
				.put("http://localhost:9901/abilities/" + ability.id, ability)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						ability.id = res.data;
						this.add(ability);
					}

					this.update=false;
				});
		},
		add: function(ability) {
			this.$emit("add", ability);
		},
		remove: function() {
			console.log("remove");
			this.$emit("remove");
		},
		getLabel: function(item) {
			if (!item) {
				return;
			}
			return item.original_name;
		},
		updateItems(text) {
			this.items = this.abilitiesList.filter(item =>
				item.original_name.toLowerCase().startsWith(text.toLowerCase())
			);
		},
		selectedItem(item) {
			console.log("SELECTED", item, this.ability);
			this.selectedAbility = item;
			// this.ability.id = item.id;
			// this.selectedAbility.name = item.name;
			// this.selectedAbility.original_name = item.original_name;
			// this.selectedAbility.type = item.type;
			// this.selectedAbility.description = item.description;
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
