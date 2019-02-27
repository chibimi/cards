<template>
	<div class="w-100">
		<div class="row">
		<div class="col-4">
			<v-autocomplete
				v-if="!ability.id"
				:items="items"
				:get-label="getLabel"
				@update-items="updateItems"
				:component-item="template"
				:auto-select-one-item="false"
				@item-selected="selectedItem"
			></v-autocomplete>

			<input v-model="selectedAbility.original_name" type="text" class="form-control" placeholder>
			<input v-model="selectedAbility.name" type="text" class="form-control" placeholder>
			
			<button v-if="ability.id || selectedAbility.id" type="submit" class="form-control btn btn-success" @click="save(selectedAbility)">Update Ability</button>
			<button v-if="ability.id" type="submit" class="form-control btn btn-danger" @click="remove(selectedAbility)">Delete Ability</button>
			<button v-if="!ability.id && selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="add(selectedAbility)">Add Ability</button>
			<button v-if="!ability.id && !selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="save(selectedAbility)">Save & Add Ability</button>
		</div>

		<div class="col-8">
			<textarea v-model="selectedAbility.description" type="text" class="form-control" rows="6" placeholder/>
		</div>
		</div>
		<hr>
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
		}
	},
	created: function() {
		this.selectedAbility = this.ability;
	},
	data() {
		return {
			selectedAbility: {},
			template: ItemTemplate,
			items: []
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
				});
		},
		add: function(ability) {
			this.$emit("add", ability);
		},
		remove: function() {
			console.log("remove")
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
			console.log("SELECTED",item, this.ability)
			this.selectedAbility = item
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
